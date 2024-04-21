package middleware

import (
	"fmt"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/services"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(config utils.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		if authorization == "" {
			ctx.String(http.StatusForbidden, "No Authorization header provided")
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		if tokenString == authorization {
			ctx.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
			ctx.Abort()
			return
		}

		token, errValidateJWT := services.ValidateJWT(tokenString, []byte(config.JwtSecretKey))
		if errValidateJWT != nil {
			log.Errorf("Error validate JWT, error %v", errValidateJWT)
			utils.Error(ctx, http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		if !token.Valid {
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Could not find bearer token in Authorization header")
			ctx.Abort()
			return
		}

		// Parse token to struct user and store into context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		utils.ShowInfoLogs(fmt.Sprintf("User login: %v", utils.LogFull(claims)))
		var currentUser models.AuthClaim
		errDecode := mapstructure.Decode(claims, &currentUser)
		if errDecode != nil {
			log.Errorf("Cannot decode claims %v to current user, error %v", utils.LogFull(claims), errDecode)
		}
		ctx.Set(utils.CurrentUser, currentUser)
		ctx.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
