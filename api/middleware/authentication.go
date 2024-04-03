package middleware

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SEC-Jobstreet/backend-application-service/utils"
	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// claims component of jwt contains many fields, we need only roles of backend-application-service
// "backend-application-service":{"backend-application-service":{"roles":["applicant","employer","admin"]}},
type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	ApplicationServiceClient clientRoles `json:"ApplicationServiceClient,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

func IsAuthorizedJWT(config *utils.Config, role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		accessToken := fields[1]

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Timeout:   time.Duration(6000) * time.Second,
			Transport: tr,
		}
		// var RealmConfigURL string = "http://10.66.29.167:9999/auth/realms/DEMOREALM"

		clientCtx := oidc.ClientContext(context.Background(), client)
		provider, err := oidc.NewProvider(clientCtx, config.KeyCloak.BaseUrl+"/auth/realms/"+config.KeyCloak.Realm)
		if err != nil {
			err := fmt.Errorf("authorisation failed while getting the provider:  %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		oidcConfig := &oidc.Config{
			ClientID: config.KeyCloak.RestApi.ClientId,
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(clientCtx, accessToken)
		if err != nil {
			err := fmt.Errorf("authorisation failed while verifying the token:  %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		var IDTokenClaims Claims // ID Token payload is just JSON.
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			err := fmt.Errorf("claims:  %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}
		fmt.Println(IDTokenClaims)
		//checking the roles
		user_access_roles := IDTokenClaims.ResourceAccess.ApplicationServiceClient.Roles
		for _, b := range user_access_roles {
			if b == role {
				// ctx.Set(authorizationPayloadKey, IDTokenClaims)
				ctx.Next()
				return
			}
		}
		err = fmt.Errorf("user not allowed to access this api:  %s", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
	}
}
