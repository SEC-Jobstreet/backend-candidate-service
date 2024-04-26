package models

import "github.com/markbates/goth"

type OAuthRequest struct {
	CurrentUrl string `form:"current_url"`
}

type OAuthResponse struct {
	goth.User
	CurrentUrl string `json:"current_url,omitempty"`
}

type OAuthGoogleRequest struct {
	CurrentUrl string `form:"current_url"`
}

type OAuthGoogleResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	CurrentUrl   string `json:"current_url,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type OAuthUserGoogleInfo struct {
	Azp              string `json:"azp,omitempty"`
	Aud              string `json:"aud,omitempty"`
	Sub              string `json:"sub,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Exp              string `json:"exp,omitempty"`
	ExpiresIn        string `json:"expires_in,omitempty"`
	Email            string `json:"email,omitempty"`
	EmailVerified    string `json:"email_verified,omitempty"`
	AccessType       string `json:"access_type,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type OAuthGoogleAccessTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	IdToken     string `json:"id_token,omitempty"`
}

type AuthClaim struct {
	Sub           string   `json:"sub"`
	CognitoGroups []string `json:"cognito:groups"`
	Iss           string   `json:"iss"`
	Version       int      `json:"version"`
	ClientId      string   `json:"client_id"`
	TokenUse      string   `json:"token_use"`
	Scope         string   `json:"scope"`
	AuthTime      int      `json:"auth_time"`
	Exp           int      `json:"exp"`
	Iat           int      `json:"iat"`
	Jti           string   `json:"jti"`
	Username      string   `json:"username"`
}
