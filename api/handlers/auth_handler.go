package handlers

import (
	"context"
	"net/http"

	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
)

type AuthHandler interface {
	HandleGoogleCallback(ctx *gin.Context)
	HandleAuthGoogle(ctx *gin.Context)
}

type authHandler struct {
	config      utils.Config
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService, config utils.Config) AuthHandler {
	return &authHandler{
		authService: authService,
		config:      config,
	}
}

func (h *authHandler) HandleGoogleCallback(ctx *gin.Context) {
	response, errHandle := h.authService.HandleGoogleCallback(ctx, h.config)
	if errHandle != nil {
		ctx.Redirect(http.StatusMovedPermanently, response.CurrentUrl)
		return
	}

	logrus.Printf("HandleGoogleCallback - response from goole callback: %v", utils.LogFull(response))
	ctx.SetCookie(utils.AccessToken, response.AccessToken, 3600, "/", "", true, false)
	ctx.SetCookie(utils.RefreshToken, response.RefreshToken, 7200, "/", "", true, false)
	ctx.SetCookie(utils.IdToken, response.IDToken, 3600, "/", "", true, false)
	ctx.SetCookie(utils.CurrentUrl, "", -1, "/", "", true, true)

	ctx.Redirect(http.StatusMovedPermanently, response.CurrentUrl)
}

func (h *authHandler) HandleAuthGoogle(ctx *gin.Context) {
	var req models.OAuthGoogleRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		logrus.Errorf("Error binding request body, err %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := ctx.Request.WithContext(context.WithValue(context.Background(), "provider", ctx.Param("provider")))

	// Store current url
	ctx.SetCookie(utils.CurrentUrl, req.CurrentUrl, 3600, "/", "", false, true)

	// Begin auth google
	gothic.BeginAuthHandler(ctx.Writer, request)
}
