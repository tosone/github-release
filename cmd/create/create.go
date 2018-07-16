package create

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

// Initialize create a new release on tag to GitHub
func Initialize(dir string, files ...string) {
	var err error
	var changeLog []byte
	var tag string
	if changeLog, tag, err = git.ChangeLog(dir); err != nil {
		logging.Fatal(err)
	}
	var releaseID uint
	if releaseID, err = release.Check(tag); err != nil {
		logging.Fatal(err)
	} else if releaseID != 0 {
		if viper.GetBool("Rewrite") {
			if err = release.Delete(releaseID); err != nil {
				logging.Fatal(err)
			}
		} else {
			logging.Info("Release exist already.")
			return
		}
	}
	var releaseReq = req.Release{
		TagName:         tag,
		TargetCommitish: viper.GetString("Branch"),
		Name:            tag,
		Body:            string(changeLog),
		Draft:           viper.GetBool("Draft"),
		Prerelease:      viper.GetBool("Prerelease"),
	}
	var releaseResp resp.Release
	if releaseResp, err = release.Create(releaseReq); err != nil {
		logging.Fatal(err)
	}
	var uploadURL = strings.TrimSuffix(releaseResp.UploadURL, common.UploadURLUseless)

	// Collect all files that will upload to the release assets.
	var releaseFiles []string
	releaseFiles = append(releaseFiles, files...)

	for _, file := range viper.GetStringSlice("Release.Files") {
		var fileList []string
		if fileList, err = filepath.Glob(file); err != nil {
			logging.Error(err)
			continue
		}
		for _, f := range fileList {
			if !com.IsFile(f) {
				logging.Error(fmt.Sprintf("No such a file: %s", f))
				continue
			}
			releaseFiles = append(releaseFiles, f)
		}
	}

	if len(releaseFiles) == 0 {
		logging.Fatal("Cannot find anything to release.")
	}

	if viper.GetBool("Release.Compress") {
		var compressFiles []string
		var compressWithSlice = viper.GetStringSlice("Release.CompressWith")
		for _, compressWith := range compressWithSlice {
			var fileList []string
			if fileList, err = filepath.Glob(compressWith); err != nil {
				logging.Error(err)
				continue
			}
			for _, f := range fileList {
				if !com.IsFile(f) {
					logging.Error(fmt.Sprintf("No such a file: %s", f))
					continue
				}
				compressFiles = append(compressFiles, f)
			}
		}
		if len(compressFiles) < len(compressWithSlice) {
			logging.Error("Some file cannot be found in `CompressWith`.")
		} else {
			for _, file := range releaseFiles {
				var filesWillCompress []string
				if !com.IsFile(file) {
					logging.Error(fmt.Sprintf("No such a file: %s", file))
					continue
				}
				filesWillCompress = append(filesWillCompress, file)
				filesWillCompress = append(filesWillCompress, compressFiles...)
				var compressPackage = path.Join(os.TempDir(), filepath.Base(file)+".tar.gz")
				if err = archiver.TarGz(compressPackage, filesWillCompress); err != nil {
					logging.Error(err)
					continue
				}
				if err = release.Upload(uploadURL, compressPackage); err != nil {
					logging.Error(err)
				}
				if com.IsFile(compressPackage) {
					if err = os.Remove(compressPackage); err != nil {
						logging.Error(err)
					}
				}
			}
		}
	} else {
		for _, file := range releaseFiles {
			if err = release.Upload(uploadURL, file); err != nil {
				logging.Error(err)
			}
		}
	}

	logging.Info(fmt.Sprintf("Release to Github successful. Please see it at %s.", releaseResp.HTMLURL))
}
