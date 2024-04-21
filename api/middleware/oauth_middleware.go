package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SEC-Jobstreet/backend-candidate-service/api/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/providers/google"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"strconv"
	"time"
)

func initOAuth2Config(config utils.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OAuthGoogleClientId,
		ClientSecret: config.OAuthGoogleClientSecret,
		RedirectURL:  config.OAuthGoogleCallbackUrl,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func OAuthMiddleware(config utils.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, _ := ctx.Cookie(utils.AccessToken)
		//idToken, _ := ctx.Cookie(utils.IdToken)

		// Validate token
		url := fmt.Sprintf("%s?access_token=%s", utils.UrlTokenInfo, accessToken)
		response, err := http.Get(url)
		if err != nil {
			logrus.Errorf("OAuthMiddleware - Error get tokeninfo, error = %v", err)
			utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Internal Server Error")
			ctx.Abort()
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			logrus.Errorf("OAuthMiddleware - Error getting token info, error = %v", utils.LogFull(response))
			utils.ErrorWithMessage(ctx, http.StatusBadRequest, "Token is invalid")
			ctx.Abort()
			return
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("OAuthMiddleware - Error reading body: %v", err)
			utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Internal Server Error")
			ctx.Abort()
			return
		}

		var oauthUserGoogle models.OAuthUserGoogleInfo
		err = json.Unmarshal(body, &oauthUserGoogle)
		if err != nil {
			logrus.Errorf("OAuthMiddleware - Error unmarshalling oauthUserGoogle info: %v", err)
			utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Internal Server Error")
			ctx.Abort()
			return
		}
		logrus.Printf("OAuthMiddleware - oauthUserGoogle = %v", utils.LogFull(oauthUserGoogle))

		if oauthUserGoogle.ErrorDescription != "" {
			logrus.Errorf("OAuthMiddleware - Error description = %v", oauthUserGoogle.ErrorDescription)
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Token is invalid")
			ctx.Abort()
			return
		}

		// Expired -> get new token
		expiresIn, _ := strconv.Atoi(oauthUserGoogle.ExpiresIn)
		if expiresIn <= 0 {
			fmt.Println("OAuthMiddleware - Token is expired")
			oauth2Config := initOAuth2Config(config)
			refreshToken, _ := ctx.Cookie(utils.RefreshToken)
			tokenSource := oauth2Config.TokenSource(context.Background(), &oauth2.Token{
				RefreshToken: refreshToken,
				Expiry:       time.Now().Add(1 * time.Hour),
			})
			newToken, err := tokenSource.Token()
			if err != nil || newToken == nil {
				logrus.Errorf("OAuthMiddleware - Error getting token from refresh token: %v", err)
				utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Could not get access token")
				ctx.Abort()
				return
			}
			ctx.SetCookie(utils.AccessToken, (*newToken).AccessToken, 3600, "/", "", true, false)
		}
		ctx.Next()
	}
}
