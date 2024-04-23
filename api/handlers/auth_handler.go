package handlers

import (
	"context"
	"net/http"

	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
)

type AuthHandler interface {
	HandleCallback(ctx *gin.Context)
	HandleAuth(ctx *gin.Context)
	HandleRefresh(ctx *gin.Context)
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

func (h *authHandler) HandleCallback(ctx *gin.Context) {
	response, errHandle := h.authService.HandleCallback(ctx, h.config)
	if errHandle != nil {
		ctx.Redirect(http.StatusMovedPermanently, response.CurrentUrl)
		return
	}

	// save user's information to DB
	// fmt.Println(response)
	// all information in response

	ctx.SetCookie(utils.AccessToken, response.AccessToken, 0, "/", "", true, false)
	ctx.SetCookie(utils.RefreshToken, response.RefreshToken, 0, "/", "", true, false)
	ctx.SetCookie(utils.IdToken, response.IDToken, 0, "/", "", true, false)
	ctx.SetCookie(utils.CurrentUrl, "", -1, "/", "", true, true)

	ctx.Redirect(http.StatusMovedPermanently, response.CurrentUrl)
}

func (h *authHandler) HandleAuth(ctx *gin.Context) {
	var req models.OAuthRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		logrus.Errorf("Error binding request body, err %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := ctx.Request.WithContext(context.WithValue(context.Background(), "provider", ctx.Param("provider")))

	// Store current url
	ctx.SetCookie(utils.CurrentUrl, req.CurrentUrl, 3600, "/", "", false, true)

	// Begin auth
	gothic.BeginAuthHandler(ctx.Writer, request)
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type refreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *authHandler) HandleRefresh(ctx *gin.Context) {
	providerName := ctx.Param("provider")
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req refreshTokenRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := provider.RefreshToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshTokenResponse := refreshTokenResponse{
		AccessToken: accessToken.AccessToken,
	}

	ctx.JSON(http.StatusOK, refreshTokenResponse)
}
