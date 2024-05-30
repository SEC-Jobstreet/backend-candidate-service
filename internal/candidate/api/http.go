package api

import "github.com/gin-gonic/gin"

type CandidateHandlers interface {
	CreateProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	GetProfileByCandidate(ctx *gin.Context)
	GetProfileByEmployer(ctx *gin.Context)
}
