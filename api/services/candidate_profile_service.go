package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	db "github.com/SEC-Jobstreet/backend-candidate-service/db/sqlc"
	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internals"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CandidateProfileService interface {
	GetProfile(ctx *gin.Context) (models.GetCandidateProfilesResponse, *models.AppError)
	GetProfileByUserId(ctx *gin.Context, userId string) (db.GetCandidateProfilesRow, *models.AppError)
	UpdateProfile(ctx *gin.Context, req models.UserProfileEditRequest) *models.AppError
	CreateProfile(req models.CandidateProfile) *models.AppError
}
type candidateProfileService struct {
	store      db.Store
	awsHandler *externals.AWSHandler
	config     utils.Config
}

func NewCandidateProfileService(store db.Store, awsHandler *externals.AWSHandler, config utils.Config) CandidateProfileService {
	return &candidateProfileService{
		store:      store,
		awsHandler: awsHandler,
		config:     config,
	}
}

func (s *candidateProfileService) initData() models.ProfileResponse {
	return models.ProfileResponse{
		ProfileUpdated: true,
		Profile: models.Profile{
			GivenName:       "Thuận",
			SurName:         "Nguyễn",
			PhoneNumber:     "+84337256835",
			CurrentLocation: "Quận 5, Thành phố Hồ Chí Minh, VN",
			PrivacySetting:  "STANDARD",
			WorkEligibility: map[string]string{"VN": "ELIGIBLE"},
			Resume:          "CV-NGUYEN-HUYNH-MINH-THUAN.pdf",
			CurrentRole:     "Developer",
			ShiftAvailability: models.ShiftAvailability{
				AnyTimeShiftAvailability: true,
				SpecificShiftAvailability: map[string]models.ShiftTimes{
					"all":       {Morning: false, Afternoon: false, Evening: false},
					"monday":    {Morning: true, Afternoon: true, Evening: true},
					"tuesday":   {Morning: true, Afternoon: true, Evening: true},
					"wednesday": {Morning: true, Afternoon: true, Evening: true},
					"thursday":  {Morning: true, Afternoon: true, Evening: true},
					"friday":    {Morning: true, Afternoon: true, Evening: true},
					"saturday":  {Morning: true, Afternoon: true, Evening: true},
					"sunday":    {Morning: true, Afternoon: true, Evening: true},
				},
			},
			CreatedAt:                  time.Now(), // Placeholder, adjust as needed
			UpdatedAt:                  time.Now(), // Placeholder, adjust as needed
			CurrentLocationCoordinates: models.Coordinates{Lat: 10.7628356, Long: 106.6824824},
			CandidateId:                "11550913",
			Email:                      "nguyenthuanit265@gmail.com",
			PhoneNumberCountryAlpha2:   "VN",
		},
		MetaTitle: "Hồ sơ cá nhân | JobStreet",
		Site: models.Site{
			ID:    "vn",
			Hosts: []string{"www.jobstreet.vn", "vn.seek.com"},
			Country: models.Country{
				IsoCode:         "VN",
				Name:            "Vietnam",
				NameWithArticle: "Vietnam",
				LanguageTags:    []string{"vi-VN"},
			},
			Brand: models.Brand{
				ID:            "JOBST",
				Name:          "JobStreet",
				CopyrightName: "Job Seeker Pty Ltd",
			},
			Analytics: models.Analytics{
				Prod:    "G-WVKCLLK2L1",
				Sandbox: "G-ZXE1FEGS1B",
			},
		},
		IndustriesAndRoles: []models.IndustryAndRole{
			{
				Value:       "other",
				DisplayName: "NA",
				Roles: []models.Role{
					{
						Value:       "other",
						DisplayName: "NA",
					},
				},
			},
		},
	}
}

func (s *candidateProfileService) GetProfile(ctx *gin.Context) (models.GetCandidateProfilesResponse, *models.AppError) {
	var response models.GetCandidateProfilesResponse
	// Get current user
	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		logrus.Errorf("GetProfile - Error get current user, error = %v ", err)
		return response, &models.AppError{
			Code:    http.StatusUnauthorized,
			Error:   err.Error,
			Message: err.Message,
		}
	}

	profiles, errGetProfile := s.store.GetCandidateProfiles(ctx, currentUser.Username)
	if errGetProfile != nil {
		logrus.Errorf("GetProfile - Error get profile, error = %v", errGetProfile)
		return response, &models.AppError{Error: errGetProfile, Code: http.StatusInternalServerError, Message: utils.INTERNAL_SERVER_ERROR}
	}

	// Mapping data
	if len(profiles) == 0 {
		return response, nil
	}

	marshal, _ := json.Marshal(profiles[0])
	json.Unmarshal(marshal, &response)
	json.Unmarshal(profiles[0].WorkEligibility, &response.WorkEligibility)
	json.Unmarshal(profiles[0].WorkShift, &response.WorkShift)
	return response, nil
}

func (s *candidateProfileService) GetProfileByUserId(ctx *gin.Context, userId string) (db.GetCandidateProfilesRow, *models.AppError) {
	profiles, errGetProfile := s.store.GetCandidateProfiles(ctx, userId)
	if errGetProfile != nil {
		logrus.Errorf("GetProfileByUserId - Error get profile, error = %v", errGetProfile)
		return db.GetCandidateProfilesRow{}, &models.AppError{Error: errGetProfile, Code: http.StatusInternalServerError, Message: utils.INTERNAL_SERVER_ERROR}
	}

	// Mapping data
	if len(profiles) == 0 {
		return db.GetCandidateProfilesRow{}, nil
	}

	return profiles[0], nil
}

func (s *candidateProfileService) UpdateProfile(ctx *gin.Context, req models.UserProfileEditRequest) *models.AppError {
	// Validate request
	validate := validator.New()
	errValidate := validate.Struct(req)
	if errValidate != nil {
		if _, ok := errValidate.(*validator.InvalidValidationError); ok {
			logrus.Errorf("UpdateProfile - errValidate = %v", errValidate)
			return &models.AppError{
				Code:    http.StatusBadRequest,
				Error:   errValidate,
				Message: errValidate.Error(),
			}
		}

		for _, err := range errValidate.(validator.ValidationErrors) {
			logrus.Errorf("UpdateProfile - Error: Field '%s' failed validation '%s'\n", err.Field(), err.Tag())
			return &models.AppError{
				Code:    http.StatusBadRequest,
				Error:   errValidate,
				Message: fmt.Sprintf("Field '%s' failed validation '%s'", err.Field(), err.Tag()),
			}
		}

		return &models.AppError{
			Code:    http.StatusBadRequest,
			Error:   errValidate,
			Message: errValidate.Error(),
		}
	}

	// Get current user
	currentUser, err := utils.GetCurrentUser(ctx)
	db := internals.GetDb()
	if err != nil {
		logrus.Errorf("UpdateProfile - Error get current user, error = %v ", err)
		return &models.AppError{
			Code:    http.StatusUnauthorized,
			Error:   err.Error,
			Message: err.Message,
		}
	}

	// Mapping data
	var reqUpdate models.CandidateProfile
	profileFromDB, _ := s.GetProfileByUserId(ctx, currentUser.Username)
	marshal, _ := json.Marshal(profileFromDB)
	json.Unmarshal(marshal, &reqUpdate)
	if len(strings.TrimSpace(req.LastName)) > 0 {
		reqUpdate.LastName = pgtype.Text{String: req.LastName, Valid: true}
	}
	if len(strings.TrimSpace(req.FirstName)) > 0 {
		reqUpdate.FirstName = pgtype.Text{String: req.FirstName, Valid: true}
	}
	if len(strings.TrimSpace(req.PhoneNumberCountry)) > 0 {
		reqUpdate.PhoneNumberCountry = pgtype.Text{String: req.PhoneNumberCountry, Valid: true}
	}
	if len(strings.TrimSpace(req.PhoneNumber)) > 0 {
		reqUpdate.Phone = pgtype.Text{String: req.PhoneNumber, Valid: true}
	}
	if len(strings.TrimSpace(req.LocationLon)) > 0 {
		locationLon, _ := strconv.ParseFloat(req.LocationLon, 10)
		reqUpdate.LocationLon = pgtype.Float8{Float64: locationLon, Valid: true}
	}
	if len(strings.TrimSpace(req.LocationLat)) > 0 {
		locationLat, _ := strconv.ParseFloat(req.LocationLat, 10)
		reqUpdate.LocationLon = pgtype.Float8{Float64: locationLat, Valid: true}
	}
	if len(strings.TrimSpace(req.CurrentRole)) > 0 {
		reqUpdate.CurrentRole = pgtype.Text{String: req.CurrentRole, Valid: true}
	}
	if len(strings.TrimSpace(req.PrivacySetting)) > 0 {
		reqUpdate.PrivacySetting = pgtype.Text{String: req.PrivacySetting, Valid: true}
	}
	workShiftByte, _ := json.Marshal(req.WorkShift)
	workEligibilityByte, _ := json.Marshal(req.WorkEligibility)
	reqUpdate.WorkShift = workShiftByte
	reqUpdate.WorkEligibility = workEligibilityByte

	// Upload resume
	file, errGetFile := ctx.FormFile("resume")
	if errGetFile != nil {
		return &models.AppError{
			Error:   errGetFile,
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: utils.RESUME_REQUIRED,
		}
	}
	originalFileName := file.Filename

	open, err2 := file.Open()
	if err2 != nil {
		logrus.Errorf("Error opening file, error = %v", err2)
		return &models.AppError{Error: err2, Code: http.StatusInternalServerError}
	}
	fileReader, _ := io.ReadAll(open)
	reader := bytes.NewReader(fileReader)
	s.awsHandler.Init(s.config)
	location, _, errUpload := s.awsHandler.Upload(originalFileName, reader)
	if errUpload != nil {
		logrus.Errorf("Error uploading file, error = %v", errUpload)
		return &models.AppError{Error: err2, Code: http.StatusInternalServerError}
	}
	reqUpdate.ResumeLink = pgtype.Text{String: location, Valid: true}

	// Trigger save user
	if len(strings.TrimSpace(profileFromDB.UserID)) == 0 {
		reqUpdate.UserID = currentUser.Username
		reqUpdate.GoogleID = pgtype.Int8{Valid: true}
		errCreate := s.CreateProfile(reqUpdate)
		if errCreate != nil {
			return &models.AppError{
				Error:   errCreate.Error,
				Code:    http.StatusInternalServerError,
				Message: errCreate.Message,
			}
		}
	}

	// Update
	db.Omit("created_at").Where("user_id = ?", reqUpdate.UserID).Updates(reqUpdate)
	return nil
}

func (s *candidateProfileService) CreateProfile(reqCreate models.CandidateProfile) *models.AppError {
	db := internals.GetDb()
	tx := db.Create(&reqCreate)
	if tx.Error != nil {
		logrus.Errorf("CreateProfile - Error create profile, error = %v", tx.Error)
		return &models.AppError{
			Error:   tx.Error,
			Code:    http.StatusInternalServerError,
			Message: utils.INTERNAL_SERVER_ERROR,
		}
	}

	return nil
}
