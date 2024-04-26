package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The value are read by viper from a config file or environment variables.
type Config struct {
	Environment             string `mapstructure:"ENVIRONMENT"`
	RESTfulServerAddress    string `mapstructure:"RESTfulServerAddress"`
	DBSource                string `mapstructure:"DB_SOURCE"`
	MigrationURL            string `mapstructure:"MIGRATION_URL"`
	FrontendURL             string `mapstructure:"FRONTEND_URL"`
	OAuthGoogleClientId     string `mapstructure:"OAUTH_GOOGLE_CLIENT_ID"`
	OAuthGoogleClientSecret string `mapstructure:"OAUTH_GOOGLE_CLIENT_SECRET"`
	OAuthGoogleCallbackUrl  string `mapstructure:"OAUTH_GOOGLE_CALLBACK_URL"`
	CognitoRegion           string `mapstructure:"COGNITO_REGION"`
	CognitoUserPoolID       string `mapstructure:"COGNITO_USER_POOL_ID"`
	S3AccessKeyId           string `mapstructure:"S3_ACCESS_KEY_ID"`
	S3SecretAccessKey       string `mapstructure:"S3_SECRET_ACCESS_KEY"`
	S3Region                string `mapstructure:"S3_REGION"`
	S3DisableSSL            bool   `mapstructure:"S3_DISABLE_SSL"`
	S3ForcePathStyle        bool   `mapstructure:"S3_FORCE_PATH_STYLE"`
	S3EndPoint              string `mapstructure:"S3_END_POINT"`
	S3BucketName            string `mapstructure:"S3_BUCKET_NAME"`
	S3BucketSubFolder       string `mapstructure:"S3_BUCKET_SUB_FOLDER"`
}

// LoadConfig reads configuration from file or environment variable.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
