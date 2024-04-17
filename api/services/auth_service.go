package services

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthService interface {
	HandleGoogleCallback(ctx *gin.Context) (models.OAuthGoogleResponse, *models.AppError)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) HandleGoogleCallback(ctx *gin.Context) (models.OAuthGoogleResponse, *models.AppError) {
	var response models.OAuthGoogleResponse
	gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("HandleGoogleCallback - Error = %v", err)
		return response, &models.AppError{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	logrus.Printf("HandleGoogleCallback - gothUser = %v\n", utils.LogFull(gothUser))
	currentURL, err := ctx.Cookie("current-url")
	if err != nil || currentURL == "" {
		currentURL = "/"
	}
	response.AccessToken = gothUser.AccessToken
	response.RefreshToken = gothUser.RefreshToken
	response.CurrentUrl = currentURL

	return response, nil
}
