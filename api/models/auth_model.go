package models

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

type AuthClaim struct {
	Id          int      `json:"id,omitempty" mapstructure:"id,omitempty"`
	FullName    string   `json:"full_name,omitempty" mapstructure:"full_name,omitempty"`
	Email       string   `json:"email,omitempty" mapstructure:"email,omitempty"`
	Roles       []string `json:"roles,omitempty" mapstructure:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty" mapstructure:"permissions,omitempty"`
	Authorized  bool     `json:"authorized,omitempty" mapstructure:"authorized,omitempty"`
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
