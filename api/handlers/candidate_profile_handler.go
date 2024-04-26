package handlers

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CandidateProfileHandler interface {
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type candidateProfileHandler struct {
	config                  utils.Config
	candidateProfileService services.CandidateProfileService
}

func NewCandidateProfileHandler(candidateProfileService services.CandidateProfileService, config utils.Config) CandidateProfileHandler {
	return &candidateProfileHandler{
		candidateProfileService: candidateProfileService,
		config:                  config,
	}
}

func (h *candidateProfileHandler) GetProfile(ctx *gin.Context) {
	res, err := h.candidateProfileService.GetProfile(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, nil)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *candidateProfileHandler) UpdateProfile(ctx *gin.Context) {
	var req models.UserProfileEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errCreate := h.candidateProfileService.UpdateProfile(ctx, req)
	if errCreate != nil {
		logrus.Errorf("ProfileHandler>CreateProfile - Failed create profile. Error = %v", errCreate.Error)
		ctx.JSON(errCreate.Code, gin.H{"error": errCreate.Message})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}
