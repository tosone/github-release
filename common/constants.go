package common

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// AppName ..
const AppName = "release2github"

// UploadURLUseless ..
const UploadURLUseless = "{?name,label}"

// APIGithub ..
const APIGithub = "https://api.github.com"

// Config ..
const Config = ".release.yml"

// EnvPrefix ..
const EnvPrefix = "release"

// RepoURL ..
func RepoURL() string {
	return fmt.Sprintf("%s/repos/%s/%s",
		APIGithub, viper.GetString("Username"), viper.GetString("Repo"))
}

// OAuthClientQueryString ..
func OAuthClientQueryString() string {
	var ClientID = viper.GetString("ClientID")
	var ClientSecret = viper.GetString("ClientSecret")
	if ClientID == "" || ClientSecret == "" {
		return ""
	}
	return fmt.Sprintf("?client_id=%s&client_secret=%s", ClientID, ClientSecret)
}

// Timeout ..
func Timeout() time.Duration {
	return time.Second * time.Duration(viper.GetInt("Runtime.Timeout"))
}

// Token ..
func Token() string {
	return fmt.Sprintf("token %s", viper.GetString("Token"))
}
