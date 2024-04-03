package services

import (
	"github.com/SEC-Jobstreet/backend-application-service/api/models"
	"github.com/SEC-Jobstreet/backend-application-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
	"time"
)

type ProfileService interface {
	GetProfile(ctx *gin.Context) (models.ProfileResponse, *models.AppError)
	CreateProfile(ctx *gin.Context, req models.UserProfileEditRequest) *models.AppError
}
type profileService struct {
}

func NewProfileService() ProfileService {
	return &profileService{}
}

func (s *profileService) GetProfile(ctx *gin.Context) (models.ProfileResponse, *models.AppError) {
	var res models.ProfileResponse
	res = s.initData()
	return res, nil
}

func (s *profileService) initData() models.ProfileResponse {
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

func (s *profileService) CreateProfile(ctx *gin.Context, req models.UserProfileEditRequest) *models.AppError {
	file, err := ctx.FormFile("resume")
	if err != nil {
		return &models.AppError{
			Error:   err,
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: utils.RESUME_REQUIRED,
		}
	}

	originalFileName := file.Filename
	dst := filepath.Join(utils.DEST_FILE_SYSTEM_RESUME, originalFileName)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		logrus.Errorf("ProfileService>CreateProfile - Could not save file: %v", err)
		return &models.AppError{
			Error:   err,
			IsError: true,
			Code:    http.StatusInternalServerError,
			Message: utils.INTERNAL_SERVER_ERROR,
		}
	}

	// TODO save into database

	return nil
}
