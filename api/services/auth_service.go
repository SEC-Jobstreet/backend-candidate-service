package services

import (
	"net/http"

	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	HandleCallback(ctx *gin.Context, config utils.Config) (models.OAuthResponse, *models.AppError)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) HandleCallback(ctx *gin.Context, config utils.Config) (models.OAuthResponse, *models.AppError) {
	var response models.OAuthResponse
	gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("HandleCallback - Error complete user auth , error = %v", err)
		return response, &models.AppError{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	currentURL, err := ctx.Cookie(utils.CurrentUrl)
	if err != nil || currentURL == "" {
		currentURL = config.FrontendURL
	}

	response.User = gothUser
	response.CurrentUrl = currentURL

	return response, nil
}
