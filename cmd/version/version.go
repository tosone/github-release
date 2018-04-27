package version

import (
	"fmt"
	"runtime"

	"github.com/tosone/release2github/common"
)

// Version build version
var Version = "no provided"

// BuildStamp build timestamp
var BuildStamp = "no provided"

// GitHash build git hash
var GitHash = "no provided"

// Setting set build info
func Setting(version, buildStamp, gitHash string) {
	Version = version
	BuildStamp = buildStamp
	GitHash = gitHash
}

// Initialize version command
func Initialize() {
	fmt.Printf("%s %s/%s\n", common.AppName, runtime.GOOS, runtime.GOARCH)
	fmt.Printf("BuildVersion: %s\n", Version)
	fmt.Printf("BuildHash: %s\n", GitHash)
	fmt.Printf("BuildDate: %s\n", BuildStamp)
}
