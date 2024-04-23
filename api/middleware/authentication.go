package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func IsAuthorizedJWT(config utils.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		accessToken := fields[1]

		// Validate token
		url := fmt.Sprintf("%s?access_token=%s", utils.UrlTokenInfo, accessToken)
		response, err := http.Get(url)

		if err != nil {
			err := fmt.Errorf("OAuthMiddleware - Error get tokeninfo, error = %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "access token invalid")
			return
		}

		ctx.Next()
	}
}
