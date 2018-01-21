package main

import (
	"fmt"
	"runtime"

	"github.com/tosone/logging"
	"github.com/tosone/release2github/cmd"
	"github.com/tosone/release2github/cmd/version"
	"github.com/tosone/release2github/common"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

func main() {
	if runtime.GOOS == "windows" {
		logging.Panic(fmt.Sprintf("%s not support windows just linux.", common.AppName))
	}

	version.Setting(Version, BuildStamp, GitHash)

	if err := cmd.RootCmd.Execute(); err != nil {
		logging.Panic(err.Error())
	}
}
