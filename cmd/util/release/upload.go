package release

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"

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
		url = fmt.Sprintf("%s&name=%s", url, filepath.Base(file))
	} else {
		url = fmt.Sprintf("%s%s&name=%s", url, common.OAuthClientQueryString(), filepath.Base(file))
	}

	var fileInfo *os.File
	if fileInfo, err = os.Open(file); err != nil {
		return
	}
	defer fileInfo.Close()
	var request *http.Request
	if request, err = http.NewRequest("POST", url, fileInfo); err != nil {
		return
	}

	request.Header.Set("Authorization", common.Token())
	request.Header.Set("Content-type", mime)
	request.Header.Set("Accept", "application/vnd.github.v3+json")

	var info os.FileInfo
	if info, err = fileInfo.Stat(); err != nil {
		return
	}
	request.ContentLength = info.Size()

	var dump []byte
	if viper.GetBool("Runtime.Debug") {
		if dump, err = httputil.DumpRequestOut(request, true); err != nil {
			return
		}
		fmt.Println(string(dump))
	}

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); err != nil {
		return
	}
	defer response.Body.Close()

	if viper.GetBool("Runtime.Debug") {
		if dump, err = httputil.DumpResponse(response, true); err != nil {
			return
		}
		fmt.Println(string(dump))
	}

	var respBody []byte
	if respBody, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	if response.StatusCode != 201 {
		err = fmt.Errorf("%s", string(respBody))
		return
	}
	return
}
