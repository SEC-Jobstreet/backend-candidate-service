package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/commands"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/queries"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/service"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type candidateHandlers struct {
	config utils.Config
	cs     *service.CandidateService
	router *gin.Engine
	media  *externals.AWSHandler
}

func NewCandidateHandlers(
	config utils.Config,
	cs *service.CandidateService,
	router *gin.Engine,
	media *externals.AWSHandler,
) *candidateHandlers {
	return &candidateHandlers{config: config, cs: cs, router: router, media: media}
}

type ProfileCreationRequest struct {
	FirstName    string `form:"first_name" binding:"required"`
	LastName     string `form:"last_name" binding:"required"`
	CountryPhone string `form:"country_phone" binding:"required"`
	Phone        string `form:"phone" binding:"required"`
	Address      string `form:"address" binding:"required"`
	Latitude     string `form:"latitude"`
	Longitude    string `form:"longitude"`
	Visa         *bool  `form:"visa" binding:"required"`
	Description  string `form:"description"`

	CurrentPosition string `form:"current_position"`
	StartDate       int64  `form:"start_date"`

	WorkWhenever *bool  `form:"work_whenever" binding:"required"`
	WorkShift    string `form:"work_shift"`

	ShareProfile *bool                `form:"share_profile" binding:"required"`
	Resume       multipart.FileHeader `form:"resume"`
}

func (ch *candidateHandlers) CreateProfile(ctx *gin.Context) {
	var form ProfileCreationRequest
	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	originalFileName := form.Resume.Filename
	open, err2 := form.Resume.Open()
	if err2 != nil {
		err := fmt.Errorf("error opening file, error = %v", err2)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	fileReader, _ := io.ReadAll(open)
	reader := bytes.NewReader(fileReader)
	location, _, errUpload := ch.media.Upload(originalFileName, reader)
	if errUpload != nil {
		err := fmt.Errorf("error uploading file, error = %v", errUpload)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	profile := models.Profile{
		Username: currentUser.Username,

		FirstName:    form.FirstName,
		LastName:     form.LastName,
		CountryPhone: form.CountryPhone,
		Phone:        form.Phone,
		Address:      form.Address,
		Latitude:     form.Latitude,
		Longitude:    form.Longitude,
		Visa:         *form.Visa,
		Description:  form.Description,

		CurrentPosition: form.CurrentPosition,
		StartDate:       form.StartDate,

		WorkWhenever: *form.WorkWhenever,
		WorkShift:    form.WorkShift,

		ShareProfile: *form.ShareProfile,
		ResumeLink:   location,
		ResumeName:   originalFileName,
	}

	command := commands.NewCreateProfileCommand(currentUser.Username, profile)
	err = ch.cs.Commands.CreateProfile.Handle(ctx, command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (ch *candidateHandlers) UpdateProfile(ctx *gin.Context) {
	var form ProfileCreationRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	profile := models.Profile{
		Username: currentUser.Username,

		FirstName:    form.FirstName,
		LastName:     form.LastName,
		CountryPhone: form.CountryPhone,
		Phone:        form.Phone,
		Address:      form.Address,
		Latitude:     form.Latitude,
		Longitude:    form.Longitude,
		Visa:         *form.Visa,
		Description:  form.Description,

		CurrentPosition: form.CurrentPosition,
		StartDate:       form.StartDate,

		WorkWhenever: *form.WorkWhenever,
		WorkShift:    form.WorkShift,

		ShareProfile: *form.ShareProfile,
	}

	if form.Resume.Size != 0 && form.Resume.Filename != "" {
		originalFileName := form.Resume.Filename
		open, err2 := form.Resume.Open()
		if err2 != nil {
			err := fmt.Errorf("error opening file, error = %v", err2)
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
		fileReader, _ := io.ReadAll(open)
		reader := bytes.NewReader(fileReader)
		location, _, errUpload := ch.media.Upload(originalFileName, reader)
		if errUpload != nil {
			err := fmt.Errorf("error uploading file, error = %v", errUpload)
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		profile.ResumeLink = location
		profile.ResumeName = originalFileName

	}

	command := commands.NewUpdateProfileCommand(currentUser.Username, profile)
	err = ch.cs.Commands.UpdateProfile.Handle(ctx, command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (ch *candidateHandlers) GetProfileByCandidate(ctx *gin.Context) {
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	query := queries.NewGetProfileByIDQuery(currentUser.Username)

	profileProjection, err := ch.cs.Queries.GetProfileByID.Handle(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profileProjection)
}

type ProfileCandidateByEmployerRequest struct {
	Username string `uri:"id" binding:"required"`
}

func (ch *candidateHandlers) GetProfileByEmployer(ctx *gin.Context) {
	var req ProfileCandidateByEmployerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	query := queries.NewGetProfileByIDQuery(req.Username)

	profileProjection, err := ch.cs.Queries.GetProfileByID.Handle(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profileProjection)
}

type JobRequest struct {
	JobID string `json:"job_id" binding:"required"`
}

func (ch *candidateHandlers) ApplyJob(ctx *gin.Context) {
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var req JobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	jobId, err := uuid.Parse(req.JobID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	application := models.Application{
		ID:          uuid.New(),
		JobID:       jobId,
		CandidateID: currentUser.Username,
		Status:      "APPLIED", // APPLIED
	}

	command := commands.NewApplyJobCommand(currentUser.Username+"-"+req.JobID, application)
	err = ch.cs.Commands.ApllyJob.Handle(ctx, command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, application)
}

func (ch *candidateHandlers) SaveJob(ctx *gin.Context) {
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var req JobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	jobId, err := uuid.Parse(req.JobID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	savedJob := models.SavedJob{
		ID:          uuid.New(),
		JobID:       jobId,
		CandidateID: currentUser.Username,
	}

	command := commands.NewSaveJobCommand(currentUser.Username+"-"+req.JobID, savedJob)
	err = ch.cs.Commands.SaveJob.Handle(ctx, command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, savedJob)
}

type JobUnsavingRequest struct {
	JobID string `json:"job_id" binding:"required"`
}

func (ch *candidateHandlers) UnsaveJob(ctx *gin.Context) {
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var req JobUnsavingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	jobId, err := uuid.Parse(req.JobID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	unsavedJob := models.SavedJob{
		JobID:       jobId,
		CandidateID: currentUser.Username,
	}

	command := commands.NewUnsaveJobCommand(currentUser.Username+"-"+req.JobID, unsavedJob)
	err = ch.cs.Commands.UnsaveJob.Handle(ctx, command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, unsavedJob)
}

func (ch *candidateHandlers) GetSavedJobList(ctx *gin.Context) {
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	query := queries.NewGetSavedJobListQuery(currentUser.Username)

	savedJobListProjection, err := ch.cs.Queries.GetSavedJobList.Handle(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobs": savedJobListProjection})
}

type listApplicationsRequest struct {
	JobID    string `form:"job_id"`
	PageID   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=10,max=20"`
}

type listApplicationsResponse struct {
	Total        int64            `json:"total"`
	PageID       int              `json:"page_id"`
	PageSize     int              `json:"page_size" `
	Applications []models.Profile `json:"applications"`
}

func (ch *candidateHandlers) GetAppliedCandidateList(ctx *gin.Context) {
	var req listApplicationsRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	jobid, err := uuid.Parse(req.JobID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	query := queries.NewGetAppliedCandidateListQuery(jobid, req.PageID, req.PageSize)

	applicationProjection, total, err := ch.cs.Queries.GetAppliedCandidateList.Handle(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res := listApplicationsResponse{
		Total:        total,
		PageID:       req.PageID,
		PageSize:     req.PageSize,
		Applications: applicationProjection,
	}

	ctx.JSON(http.StatusOK, res)
}

type JobIdURIRequest struct {
	JobID string `uri:"job_id" binding:"required"`
}

func (ch *candidateHandlers) GetAppliedCandidateNumber(ctx *gin.Context) {
	var req JobIdURIRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	jobid, err := uuid.Parse(req.JobID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	query := queries.NewGetAppliedCandidateNumberQuery(jobid)

	applicationProjection, err := ch.cs.Queries.GetAppliedCandidateNumber.Handle(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"total": applicationProjection})
}
