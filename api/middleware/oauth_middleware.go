package middleware

import (
	"bytes"
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
		//accessToken, _ := ctx.Cookie(utils.AccessToken)
		//idToken, _ := ctx.Cookie(utils.IdToken)

		// Validate token
		url := fmt.Sprintf("%s?access_token=%s", utils.UrlTokenInfo, "ya29.a0Ad52N38xQlLmCfMuyKD6G6NLPqRR-7QJjToEYr-wtsdZZHRMufpVHOAa9mbejJ-E-bMh_qFpcl-I6IBQA-V3cZUkRHQWl19dfCmdtbvPqy45eixjHAinh88AiLR3APZ_Ow2hDG7frH1xdqM3XqQhrFB15rt7usHqD-puaCgYKAX4SARISFQHGX2MiF5QfjBfgmOz-zeUd5Ib0vQ0171")
		response, err := http.Get(url)
		if err != nil {
			logrus.Errorf("OAuthMiddleware - Error get tokeninfo, error = %v", err)
			utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Internal Server Error")
			ctx.Abort()
			return
		}
		defer response.Body.Close()

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

		// Expired -> get new token
		expiresIn, _ := strconv.Atoi(oauthUserGoogle.ExpiresIn)
		if expiresIn <= 0 || oauthUserGoogle.ErrorDescription != "" {
			logrus.Println("OAuthMiddleware - Token is expired")
			refreshToken, _ := ctx.Cookie(utils.RefreshToken)
			newToken, err := getAccessToken(refreshToken, config)
			if err != nil || newToken == nil || (*newToken).AccessToken == "" {
				logrus.Errorf("OAuthMiddleware - Error getting token from refresh token: %v", err)
				utils.ErrorWithMessage(ctx, http.StatusInternalServerError, "Could not get access token")
				ctx.Abort()
				return
			}
			logrus.Printf("OAuthMiddleware - new access token = %v", (*newToken).AccessToken)
			ctx.SetCookie(utils.AccessToken, (*newToken).AccessToken, 3600, "/", "", true, false)
		}
		ctx.Next()
	}
}

func getAccessToken(refreshToken string, config utils.Config) (*models.OAuthGoogleAccessTokenResponse, error) {
	values := map[string]string{
		"client_id":     config.OAuthGoogleClientId,
		"client_secret": config.OAuthGoogleClientSecret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", utils.UrlGetAccessTokenV4, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse models.OAuthGoogleAccessTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
