package api

import (
	"time"

	"github.com/SEC-Jobstreet/backend-candidate-service/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := router.Group("/api/v1")

	authRoutes.POST("/create_profile", middleware.AuthMiddleware(s.config, []string{"candidates"}), s.CreateProfile)
	authRoutes.PUT("/update_profile", middleware.AuthMiddleware(s.config, []string{"candidates"}), s.UpdateProfile)
	authRoutes.GET("/profile", middleware.AuthMiddleware(s.config, []string{"candidates"}), s.GetProfileByCandidate)
	authRoutes.GET("/profile_by_employer/:id", middleware.AuthMiddleware(s.config, []string{"employers"}), s.GetProfileByEmployer)

	s.router = router
}
