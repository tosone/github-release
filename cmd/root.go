package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/unknwon/com"

	"github.com/tosone/release2github/cmd/create"
	"github.com/tosone/release2github/cmd/delete"
	"github.com/tosone/release2github/cmd/version"
	"github.com/tosone/release2github/common"
)

var dir string

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   common.AppName,
	Short: "Release files and changelog to github release page.",
	Long:  `Release files and changelog to github release page.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version",
	Long:  `The version that build detail information.`,
	Run: func(_ *cobra.Command, _ []string) {
		version.Initialize()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new release on github release page.",
	Long: `Create a new release on github release page. 
The config file is current directory .release file and something else in env about users key info.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		initConfig()
		create.Initialize(dir, args...)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tag release from github release page.",
	Long:  `Delete a tag release from github release page.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		initConfig()
		delete.Initialize(args...)
	},
}

func init() {
	var err error
	var currPath string
	if currPath, err = os.Getwd(); err != nil {
		logging.Fatal(err)
	}
	createCmd.PersistentFlags().StringVarP(&dir, "dir", "d", currPath, "execute path")

	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(deleteCmd)
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(common.EnvPrefix)
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
