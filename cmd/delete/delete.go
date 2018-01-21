package delete

import (
	"fmt"

	"github.com/tosone/logging"
	"github.com/tosone/release2github/cmd/util/release"
)

func Initialize(tag string) {
	var err error
	var releaseID uint
	if releaseID, err = release.Check(tag); err != nil {
		logging.Fatal(err)
	} else if releaseID != 0 {
		if err = release.Delete(releaseID); err != nil {
			logging.Fatal(err)
		}
		logging.Info(fmt.Sprintf("Delete %s from Github successful.", tag))
	}
}
