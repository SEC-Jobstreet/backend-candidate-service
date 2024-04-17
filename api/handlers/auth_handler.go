package handlers

import (
	"context"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthHandler interface {
	HandleGoogleCallback(ctx *gin.Context)
	HandleAuthGoogle(ctx *gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) HandleGoogleCallback(ctx *gin.Context) {
	response, errHandle := h.authService.HandleGoogleCallback(ctx)
	if errHandle != nil {
		ctx.JSON(errHandle.Code, gin.H{"error": errHandle.Message})
		return
	}

	ctx.JSON(http.StatusOK, response)

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
