package handlers

import (
	"github.com/SEC-Jobstreet/backend-application-service/api/models"
	"github.com/SEC-Jobstreet/backend-application-service/api/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProfileHandler interface {
	GetProfile(ctx *gin.Context)
	CreateProfile(ctx *gin.Context)
}

type profileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) ProfileHandler {
	return &profileHandler{
		profileService: profileService,
	}
}

func (h *profileHandler) GetProfile(ctx *gin.Context) {
	// TODO Get user login

	// TODO Get profile by email

	res, _ := h.profileService.GetProfile(ctx)
	ctx.JSON(http.StatusOK, res)
}

func (h *profileHandler) CreateProfile(ctx *gin.Context) {
	var req models.UserProfileEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errCreate := h.profileService.CreateProfile(ctx, req)
	if errCreate != nil {
		logrus.Errorf("ProfileHandler>CreateProfile - Failed create profile. Error = %v", errCreate.Error)
		ctx.JSON(errCreate.Code, gin.H{"error": errCreate.Message})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
