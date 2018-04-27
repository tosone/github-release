package release

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
	"github.com/tosone/release2github/common"
	"github.com/tosone/release2github/common/req"
	"github.com/tosone/release2github/common/resp"
)

// Create create release on GitGub
func Create(release req.Release) (releaseResp resp.Release, err error) {
	var url = fmt.Sprintf("%s/releases%s",
		common.RepoURL(), common.OAuthClientQueryString())
	response, body, errs := gorequest.New().
		Timeout(common.Timeout()).
		SetDebug(viper.GetBool("Runtime.Debug")).
		Post(url).
		Type("json").
		Set("Authorization", common.Token()).
		Set("Accept", "application/vnd.github.v3+json").
		SendStruct(release).
		End()
	if len(errs) != 0 {
		err = errs[len(errs)-1]
		return
	}
	if response.StatusCode != 201 {
		err = fmt.Errorf("%s", body)
		return
	}
	if err = json.Unmarshal([]byte(body), &releaseResp); err != nil {
		return
	}
	return
}
