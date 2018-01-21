package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Unknwon/com"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/tosone/release2github/cmd/create"
	"github.com/tosone/release2github/cmd/version"
	"github.com/tosone/release2github/common"
)

var dir string

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   common.AppName,
	Short: "",
	Long:  ``,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		version.Initialize()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		initConfig()
		fmt.Println(viper.GetStringSlice("Release.Files"))
		create.Initialize(dir, args...)
	},
}

func init() {
	var err error
	var currPath string
	if currPath, err = os.Getwd(); err != nil {
		logging.Fatal(err)
	}
	createCmd.PersistentFlags().StringVarP(&dir, "dir", "d", currPath, "config file")

	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(versionCmd)
}

func initConfig() {
	defaultConfig()
	if dir != "" {
		var config = path.Join(dir, common.Config)
		if !com.IsFile(config) {
			logging.Fatal(fmt.Sprintf("Cannot find config file here: %s", config))
		} else {
			viper.SetConfigFile(config)
		}
	} else {
		logging.Fatal("Cannot find config file. Please check.")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logging.Panic("Cannot find the special config file.")
	}
}

func defaultConfig() {
	viper.SetDefault("DatabaseEngine", "sqlite3")
	viper.SetDefault("log", "err.log")
}
