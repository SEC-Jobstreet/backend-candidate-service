package middleware

import (
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
	"strings"
)

func initOAuth2Config(config utils.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OAuthGoogleClientId,     // Replace with your client ID
		ClientSecret: config.OAuthGoogleClientSecret, // Replace with your client secret
		RedirectURL:  config.OAuthGoogleCallbackUrl,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func OAuthMiddleware(config utils.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//oauth2Config := initOAuth2Config(config)
		//refreshToken := ctx.Request.Header.Get("Refresh-Token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "No Authorization header provided")
			ctx.Abort()
			return
		}

		accessToken := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if accessToken == authorizationHeader {
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Could not find bearer token in Authorization header")
			ctx.Abort()
			return
		}

		// Validate token
		url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=%s", accessToken)
		response, err := http.Get(url)
		if err != nil {
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Token is valid")
			ctx.Abort()
			return
		}
		defer response.Body.Close()
		if response.StatusCode != 200 {
			logrus.Errorf("Error getting token info")
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Token is valid")
			ctx.Abort()
			return
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("Error reading body: %v", err)
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Token is valid")
			ctx.Abort()
			return
		}

		var oauthUserGoogle models.OAuthUserGoogleInfo
		err = json.Unmarshal(body, &oauthUserGoogle)
		if err != nil {
			logrus.Errorf("Error unmarshalling oauthUserGoogle info: %v", err)
			utils.ErrorWithMessage(ctx, http.StatusUnauthorized, "Token is valid")
			ctx.Abort()
			return
		}
		logrus.Printf("oauthUserGoogle = %v", utils.LogFull(oauthUserGoogle))

		// Get new token from refresh token
		//tokenSource := oauth2Config.TokenSource(context.Background(), &oauth2.Token{
		//	RefreshToken: refreshToken,
		//	AccessToken:  accessToken,
		//	Expiry:       time.Now().Add(1 * time.Hour),
		//})

		ctx.Next()
	}
}
