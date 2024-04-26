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
		utils.Error(ctx, http.StatusOK)
		return
	}
	utils.Ok(ctx, res)
}

func (h *candidateProfileHandler) UpdateProfile(ctx *gin.Context) {
	var req models.UserProfileEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.ErrorWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	errSave := h.candidateProfileService.UpdateProfile(ctx, req)
	if errSave != nil {
		logrus.Errorf("ProfileHandler>CreateProfile - Failed create/update profile. Error = %v", errSave.Error)
		utils.ErrorWithMessage(ctx, errSave.Code, errSave.Message)
		return
	}

	utils.Ok(ctx, req)
}
