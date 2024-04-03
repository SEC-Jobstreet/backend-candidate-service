package server

import (
	"net/http"

	"github.com/SEC-Jobstreet/backend-application-service/utils"
	"github.com/gin-gonic/gin"
)

type applyRequest struct {
	Username     string `json:"username" binding:"required,alphanum,min=3,max=15"`
	Password     string `json:"password" binding:"required,min=6"`
	FirstName    string `json:"first_name" binding:"required,min=1,max=30"`
	LastName     string `json:"last_name" binding:"required,min=1,max=30"`
	Email        string `json:"email" binding:"required,email"`
	MobileNumber string `json:"phone"`
}

type applyResponse struct {
	Status bool `json:"status"`
}

func (s *Server) Apply(ctx *gin.Context) {
	var request applyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res := &applyResponse{
		Status: true,
	}
	ctx.JSON(http.StatusOK, res)
}
