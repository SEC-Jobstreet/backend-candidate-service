package middleware

import (
	"context"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
	"net/http"
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
		oauth2Config := initOAuth2Config(config)
		token, err := ctx.Cookie("auth-token")
		if err != nil || token == "" {
			state := "state"
			url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
			ctx.Redirect(http.StatusTemporaryRedirect, url)
			ctx.Abort()
			return
		}

		tokenSource := oauth2Config.TokenSource(context.Background(), &oauth2.Token{AccessToken: token})
		newToken, err := tokenSource.Token()
		if err != nil || newToken == nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh access token"})
			return
		}

		ctx.Set("userToken", newToken) // Set the token in the context for further use
		ctx.Next()
	}
}
