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
	HandleGoogleCallback(ctx *gin.Context, config utils.Config) (models.OAuthGoogleResponse, *models.AppError)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) HandleGoogleCallback(ctx *gin.Context, config utils.Config) (models.OAuthGoogleResponse, *models.AppError) {
	var response models.OAuthGoogleResponse
	gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("HandleGoogleCallback - Error complete user auth google, error = %v", err)
		return response, &models.AppError{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	currentURL, err := ctx.Cookie(utils.CurrentUrl)
	if err != nil || currentURL == "" {
		currentURL = config.FrontendURL
	}

	response.AccessToken = gothUser.AccessToken
	response.RefreshToken = gothUser.RefreshToken
	response.IDToken = gothUser.IDToken
	response.CurrentUrl = currentURL

	return response, nil
}
