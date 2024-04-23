package models

import "github.com/markbates/goth"

type OAuthRequest struct {
	CurrentUrl string `form:"current_url"`
}

type OAuthResponse struct {
	goth.User
	CurrentUrl string `json:"current_url,omitempty"`
}
