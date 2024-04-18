package middleware

import (
	"context"
	"net/http"

	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
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
		// cai nay k hỉu
		// lấy access_token sau đó verify thử ok k. nếu hết hạn thì lấy lại access_token từ refresh_token.
		// nếu hết hạn nữa thì cancel để ng dùng đăng nhập lại.
		token, err := ctx.Cookie("auth-token") // auth-token là access_token hả
		if err != nil || token == "" {
			state := "state"
			url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline) // cần k hay trỏ thẳng đến homepage lun
			ctx.Redirect(http.StatusMovedPermanently, url)
			ctx.Abort()
			return
		}

		// sao k truyền thêm refresh token nữa
		tokenSource := oauth2Config.TokenSource(context.Background(), &oauth2.Token{AccessToken: token})
		newToken, err := tokenSource.Token()
		if err != nil || newToken == nil {
			ctx.Redirect(http.StatusMovedPermanently, config.FrontendURL)
			return
		}

		// nếu có thay đổi thì phải lưu lại vào cookie các token vừa thay đổi để request sau còn sài nựa

		// đó là suy nghĩ của t thôi m thấy sai thì nhắn t

		// ctx.Set("userToken", newToken) // Set the token in the context for further use
		ctx.Next()
	}
}
