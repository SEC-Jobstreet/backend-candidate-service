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

	//ctx.SetCookie("access_token", response.AccessToken, 0, "/", "", false, false)
	//ctx.SetCookie("refresh_token", response.RefreshToken, 0, "/", "", false, false)
	//ctx.SetCookie("IDToken", response.IDToken, 0, "/", "", false, false)
	//ctx.SetCookie("current-url", "", -1, "/", "", false, true)

	//ctx.Redirect(http.StatusMovedPermanently, response.CurrentUrl)

	utils.Ok(ctx, response)
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
	ctx.SetCookie("current-url", req.CurrentUrl, 3600, "/", "", false, true)

	// Begin auth google
	gothic.BeginAuthHandler(ctx.Writer, request)
}
