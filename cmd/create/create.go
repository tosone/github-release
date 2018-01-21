package create

import (
	"fmt"
	"strings"

	"github.com/Unknwon/com"
	"github.com/mholt/archiver"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/tosone/release2github/cmd/util/git"
	"github.com/tosone/release2github/cmd/util/release"
	"github.com/tosone/release2github/common"
	"github.com/tosone/release2github/common/req"
	"github.com/tosone/release2github/common/resp"
)

func Initialize(dir string, files ...string) {
	var err error
	var changeLog []byte
	var tag string
	if changeLog, tag, err = git.ChangeLog(dir); err != nil {
		logging.Fatal(err)
	}
	if viper.GetBool("Common.Rewrite") {
		var releaseID uint
		if releaseID, err = release.Check(tag); err != nil {
			logging.Fatal(err)
		} else if releaseID != 0 {
			if err = release.Delete(releaseID); err != nil {
				logging.Fatal(err)
			}
		}
	}
	var releaseReq = req.Release{
		TagName:         tag,
		TargetCommitish: viper.GetString("Common.Branch"),
		Name:            tag,
		Body:            string(changeLog),
		Draft:           viper.GetBool("Common.Draft"),
		Prerelease:      viper.GetBool("Common.Prerelease"),
	}
	var releaseResp resp.Release
	if releaseResp, err = release.Create(releaseReq); err != nil {
		logging.Fatal(err)
	}
	var uploadUrl = strings.TrimSuffix(releaseResp.UploadUrl, common.UploadUrlUseless)

	// Collect all files that will upload to the release assets.
	var releaseFiles []string
	releaseFiles = append(releaseFiles, files...)
	releaseFiles = append(releaseFiles, viper.GetStringSlice("Release.Files")...)

	if viper.GetBool("Release.Compress") {
		var compressFiles []string
		var compressWithSlice = viper.GetStringSlice("Release.CompressWith")
		for _, commpressWith := range compressWithSlice {
			if !com.IsFile(commpressWith) {
				logging.Error(fmt.Sprintf("No such a file: %s", commpressWith))
				break
			}
			compressFiles = append(compressFiles, commpressWith)
		}
		if len(compressFiles) != len(compressWithSlice) {
			logging.Error("Something error occured will not upload assets.")
		} else {
			for _, file := range releaseFiles {
				var filesWillCompress []string
				if !com.IsFile(file) {
					logging.Error(fmt.Sprintf("No such a file: %s", file))
					continue
				}
				filesWillCompress = append(filesWillCompress, file)
				filesWillCompress = append(filesWillCompress, compressFiles...)
				archiver.TarGz.Make(file+".tar.gz", filesWillCompress)
				if err = release.Upload(uploadUrl, file); err != nil {
					logging.Error(err)
				}
			}
		}
	} else {
		for _, file := range files {
			if err = release.Upload(uploadUrl, file); err != nil {
				logging.Error(err)
			}
		}
	}

	logging.Info(fmt.Sprintf("Release to Github successful. Please see it at %s.\n", releaseResp.HtmlUrl))
}
