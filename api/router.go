package server

import (
	"time"

	"github.com/SEC-Jobstreet/backend-application-service/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) setupRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "content-type", "accept", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := router.Group("/api/v1").Use(middleware.IsAuthorizedJWT(&s.config, "applicant"))
	authRoutes.POST("/apply_job", s.Apply)

	// Profile
	apiJob := router.Group("/internal/api")
	{
		apiJob.GET("/profiles", s.profileHandler.GetProfile)
		apiJob.POST("/profiles", s.profileHandler.CreateProfile)
	}

	s.router = router
}
