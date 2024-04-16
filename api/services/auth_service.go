package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthService interface {
	HandleGoogleCallback(ctx *gin.Context)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) HandleGoogleCallback(ctx *gin.Context) {
	gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Printf("Error = %v", err)
		return
	}

	fmt.Printf("gothUser = %s\n", gothUser.UserID)
	fmt.Printf("gothUser = %s\n", gothUser.Name)
	fmt.Printf("gothUser = %s\n", gothUser.Email)
}
