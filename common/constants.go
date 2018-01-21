package common

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const AppName = "release2github"

const UploadUrlUseless = "{?name,label}"

const ApiGithub = "https://api.github.com"

const Config = ".release"

const EnvPrefix = "release"

func RepoUrl() string {
	return fmt.Sprintf("%s/repos/%s/%s", ApiGithub, viper.GetString("Username"), viper.GetString("Repo"))
}

func OAuthClientQueryString() string {
	var ClientID = viper.GetString("ClientID")
	var ClientSecret = viper.GetString("ClientSecret")
	if ClientID == "" || ClientSecret == "" {
		return ""
	}
	return fmt.Sprintf("?client_id=%s&client_secret=%s", ClientID, ClientSecret)
}

func Timeout() time.Duration {
	return time.Second * time.Duration(viper.GetInt("Runtime.Timeout"))
}

func Token() string {
	return fmt.Sprintf("token %s", viper.GetString("Token"))
}
