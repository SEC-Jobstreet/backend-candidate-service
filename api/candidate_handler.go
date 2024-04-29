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

func (s *Server) example(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "OK")
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
	Resume       multipart.FileHeader `form:"resume" binding:"required"`
}

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

type ProfileUpdateRequest struct {
	FirstName    string `form:"first_name"`
	LastName     string `form:"last_name"`
	CountryPhone string `form:"country_phone"`
	Phone        string `form:"phone"`
	Address      string `form:"address"`
	Latitude     string `form:"latitude"`
	Longitude    string `form:"longitude"`
	Visa         *bool  `form:"visa"`
	Description  string `form:"description"`

	CurrentPosition string `form:"current_position"`
	StartDate       int64  `form:"start_date"`

	WorkWhenever *bool  `form:"work_whenever"`
	WorkShift    string `form:"work_shift"`

	ShareProfile *bool                `form:"share_profile"`
	Resume       multipart.FileHeader `form:"resume"`
}

func (s *Server) UpdateProfile(ctx *gin.Context) {
	var form ProfileUpdateRequest
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

	if len(strings.TrimSpace(form.FirstName)) > 0 {
		profile["first_name"] = form.FirstName
	}
	if len(strings.TrimSpace(form.LastName)) > 0 {
		profile["last_name"] = form.LastName
	}
	if len(strings.TrimSpace(form.CountryPhone)) > 0 {
		profile["country_phone"] = form.CountryPhone
	}
	if len(strings.TrimSpace(form.Phone)) > 0 {
		profile["phone"] = form.Phone
	}
	if len(strings.TrimSpace(form.Address)) > 0 {
		profile["address"] = form.Address
	}
	if len(strings.TrimSpace(form.Latitude)) > 0 {
		profile["latitude"] = form.Latitude
	}
	if len(strings.TrimSpace(form.Longitude)) > 0 {
		profile["longitude"] = form.Longitude
	}

	if form.Visa != nil {
		profile["visa"] = *form.Visa
	}
	if len(strings.TrimSpace(form.Description)) > 0 {
		profile["description"] = form.Description
	}
	if len(strings.TrimSpace(form.CurrentPosition)) > 0 {
		profile["current_position"] = form.CurrentPosition
	}
	if form.StartDate != 0 {
		profile["start_date"] = form.StartDate
	}
	if form.WorkWhenever != nil {
		profile["work_whenever"] = *form.WorkWhenever
	}
	profile["work_shift"] = form.WorkShift

	if form.ShareProfile != nil {
		profile["share_profile"] = *form.ShareProfile
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
