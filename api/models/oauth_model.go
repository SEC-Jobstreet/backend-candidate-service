package models

type OAuthGoogleRequest struct {
	CurrentUrl string `form:"current_url"`
}

type OAuthGoogleResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	CurrentUrl   string `json:"current_url,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"IDToken,omitempty"`
}
