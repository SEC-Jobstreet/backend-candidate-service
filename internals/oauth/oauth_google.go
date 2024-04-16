package oauth

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomKey"
	MaxAge = 86400 * 30
	IsProd = true
)

type OAuthGoogleService interface {
	NewGoogleOAuth(config utils.Config)
}

type oauthGoogleService struct {
}

func NewOAuthGoogleService() OAuthGoogleService {
	return &oauthGoogleService{}
}

func (s *oauthGoogleService) NewGoogleOAuth(config utils.Config) {
	// Config clientId, clientSecret, google callback url
	googleClientId := config.OAuthGoogleClientId
	googleClientSecret := config.OAuthGoogleClientSecret
	googleCallbackUrl := config.OAuthGoogleCallbackUrl

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, googleCallbackUrl, "email", "profile"),
	)
}
