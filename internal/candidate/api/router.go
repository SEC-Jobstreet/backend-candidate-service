package api

import (
	"time"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ch *candidateHandlers) SetupRouter() *gin.Engine {

	ch.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "content-type", "accept", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ch.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	groupRoutes := ch.router.Group("/api/v1")

	groupRoutes.POST("/create_profile", middleware.AuthMiddleware(ch.config, []string{"candidates"}), ch.CreateProfile)
	groupRoutes.PUT("/update_profile", middleware.AuthMiddleware(ch.config, []string{"candidates"}), ch.UpdateProfile)
	groupRoutes.GET("/profile", middleware.AuthMiddleware(ch.config, []string{"candidates"}), ch.GetProfileByCandidate)
	groupRoutes.GET("/profile_by_employer/:id", middleware.AuthMiddleware(ch.config, []string{"employers"}), ch.GetProfileByEmployer)

	return ch.router
}
