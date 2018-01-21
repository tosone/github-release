package release

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
	"github.com/tosone/release2github/common"
	"gopkg.in/h2non/filetype.v1"
)

func Upload(url, file string) (err error) {
	var ext = filepath.Ext(file)
	var mime string
	if ext == "" || filetype.GetType(strings.TrimPrefix(ext, ".")).Extension == "unknown" {
		mime = "application/octet-stream"
	} else {
		mime = filetype.GetType(strings.TrimPrefix(ext, ".")).MIME.Value
	}
	if strings.IndexAny(url, "?") != -1 {
		url = fmt.Sprintf("%s&name=%s", url, file)
	} else {
		url = fmt.Sprintf("%s%s&name=%s", url, common.OAuthClientQueryString(), file)
	}
	response, body, errs := gorequest.New().
		SetDebug(true).
		Timeout(common.Timeout()).
		Post(url).
		Set("Content-Type", mime).
		Set("Authorization", viper.GetString("Token")).
		Set("Accept", "application/vnd.github.v3+json").
		Type("multipart").
		SendFile(file).
		End()
	if len(errs) != 0 {
		err = errs[len(errs)-1]
		return
	}
	if response.StatusCode != 201 {
		err = fmt.Errorf("%s", body)
		return
	}
	return
}
