package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/SEC-Jobstreet/backend-candidate-service/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	ShareProfile *bool `form:"share_profile" binding:"required"`
	// swagger:ignore
	Resume multipart.FileHeader `form:"resume"`
}

// @Summary Create profile
// @Description create profile for a specific candidate
// @Tags candidates
// @Accept  multipart/form-data
// @Produce  json
// @Param candidate_id path int true "Candidate ID"
// @Param profile formData ProfileCreationRequest true "Profile data"
// @Success 200 {object} models.Candidates
// @Router /create_profile [post]
func (s *Server) CreateProfile(ctx *gin.Context) {
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
	location, _, errUpload := s.media.Upload(originalFileName, reader)
	if errUpload != nil {
		err := fmt.Errorf("error uploading file, error = %v", errUpload)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	profile := models.Candidates{
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

	err = s.store.Create(&profile).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// @Summary Update profile
// @Description update profile for a specific candidate
// @Tags candidates
// @Accept  json
// @Produce  json
// @Param candidate_id path int true "Candidate ID"
// @Param profile body ProfileCreationRequest true "Profile data"
// @Success 200 {object} models.Candidates
// @Router /update_profile [put]
func (s *Server) UpdateProfile(ctx *gin.Context) {
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

	profile := map[string]interface{}{}

	profile["first_name"] = form.FirstName
	profile["last_name"] = form.LastName
	profile["country_phone"] = form.CountryPhone
	profile["phone"] = form.Phone
	profile["address"] = form.Address

	if len(strings.TrimSpace(form.Latitude)) > 0 {
		profile["latitude"] = form.Latitude
	}
	if len(strings.TrimSpace(form.Longitude)) > 0 {
		profile["longitude"] = form.Longitude
	}

	profile["visa"] = *form.Visa
	if len(strings.TrimSpace(form.Description)) > 0 {
		profile["description"] = form.Description
	}
	if len(strings.TrimSpace(form.CurrentPosition)) > 0 {
		profile["current_position"] = form.CurrentPosition
	}
	if form.StartDate != 0 {
		profile["start_date"] = form.StartDate
	}
	profile["work_whenever"] = *form.WorkWhenever
	profile["work_shift"] = form.WorkShift

	profile["share_profile"] = *form.ShareProfile

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
		location, _, errUpload := s.media.Upload(originalFileName, reader)
		if errUpload != nil {
			err := fmt.Errorf("error uploading file, error = %v", errUpload)
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		profile["resume_name"] = originalFileName
		profile["resume_link"] = location

	}

	err = s.store.Model(&models.Candidates{Username: currentUser.Username}).Updates(profile).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// @Summary Get profile by candidate
// @Description get profile details for a specific candidate
// @Tags candidates
// @Accept  json
// @Produce  json
// @Param candidate_id path int true "Candidate ID"
// @Success 200 {object} models.Candidates
// @Router /candidates/{candidate_id}/profile [get]
func (s *Server) GetProfileByCandidate(ctx *gin.Context) {

	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	candidate := &models.Candidates{
		Username: currentUser.Username,
	}
	err = s.store.First(candidate).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, candidate)
}

type ProfileCandidateByEmployerRequest struct {
	Username string `uri:"id" binding:"required"`
}

// @Summary Get profile by employer
// @Description get profile details for a specific employer
// @Tags employers
// @Accept  json
// @Produce  json
// @Param employer_id path int true "Employer ID"
// @Success 200 {object} models.Candidates
// @Router /employers/{employer_id}/profile [get]
func (s *Server) GetProfileByEmployer(ctx *gin.Context) {

	var req ProfileCandidateByEmployerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	candidate := &models.Candidates{
		Username: req.Username,
	}
	err := s.store.First(candidate).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, candidate)
}
