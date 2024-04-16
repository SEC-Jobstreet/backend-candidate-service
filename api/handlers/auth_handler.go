package handlers

import (
	"context"
	"fmt"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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
	h.authService.HandleGoogleCallback(ctx)
}

func (h *authHandler) HandleAuthGoogle(ctx *gin.Context) {
	fmt.Printf("provider = %v", ctx.Param("provider"))
	request := ctx.Request.WithContext(context.WithValue(context.Background(), "provider", ctx.Param("provider")))
	gothic.BeginAuthHandler(ctx.Writer, request)
}
