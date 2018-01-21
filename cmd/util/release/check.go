package release

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/tosone/release2github/common"
	"github.com/tosone/release2github/common/resp"
)

func Check(tag string) (releaseID uint, err error) {
	if tag == "" {
		return
	}
	var url = fmt.Sprintf("%s/releases/tags/%s%s", common.RepoUrl(), tag, common.OAuthClientQueryString())
	response, body, errs := gorequest.New().
		Timeout(common.Timeout()).
		Get(url).
		End()
	if len(errs) != 0 {
		err = errs[len(errs)-1]
		return
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("%s", body)
		return
	}
	var release = new(resp.Release)
	if err = json.Unmarshal([]byte(body), release); err != nil {
		return
	}
	if release.TagName == tag {
		releaseID = release.ID
		return
	}
	return
}
