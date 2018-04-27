package delete

import (
	"fmt"

	"github.com/tosone/logging"
	"github.com/tosone/release2github/cmd/util/release"
)

// Initialize ..
func Initialize(tags ...string) {
	var err error
	if len(tags) == 0 {
		logging.Error("No tag release to delete.")
		return
	}
	for _, tag := range tags {
		var releaseID uint
		if releaseID, err = release.Check(tag); err != nil {
			logging.Fatal(err)
		} else if releaseID != 0 {
			if err = release.Delete(releaseID); err != nil {
				logging.Fatal(err)
			}
			logging.Info(fmt.Sprintf("Delete %s release from Github successful.", tag))
		}
	}
}
