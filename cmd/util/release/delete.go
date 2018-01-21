package release

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
	"github.com/tosone/release2github/common"
)

func Delete(releaseID uint) (err error) {
	if releaseID == 0 {
		err = fmt.Errorf("release id is not correct: %d", releaseID)
	}
	var url = fmt.Sprintf("%s/releases/%d%s", common.RepoUrl(), releaseID, common.OAuthClientQueryString())
	response, body, errs := gorequest.New().
		Timeout(common.Timeout()).
		SetDebug(viper.GetBool("Runtime.Debug")).
		CustomMethod(gorequest.DELETE, url).
		Set("Authorization", common.Token()).
		End()
	if len(errs) != 0 {
		err = errs[len(errs)-1]
		return
	}
	if response.StatusCode != 204 {
		err = fmt.Errorf("%s", body)
		return
	}
	return
}
