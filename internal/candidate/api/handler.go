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

// type ProfileUpdateRequest struct {
// 	FirstName    string `form:"first_name"`
// 	LastName     string `form:"last_name"`
// 	CountryPhone string `form:"country_phone"`
// 	Phone        string `form:"phone"`
// 	Address      string `form:"address"`
// 	Latitude     string `form:"latitude"`
// 	Longitude    string `form:"longitude"`
// 	Visa         *bool  `form:"visa"`
// 	Description  string `form:"description"`

// 	CurrentPosition string `form:"current_position"`
// 	StartDate       int64  `form:"start_date"`

// 	WorkWhenever *bool  `form:"work_whenever"`
// 	WorkShift    string `form:"work_shift"`

// 	ShareProfile *bool                `form:"share_profile"`
// 	Resume       multipart.FileHeader `form:"resume"`
// }

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
