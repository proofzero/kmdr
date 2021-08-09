/*
Copyright © 2021 kubelt

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	apply "github.com/proofzero/kmdr/cmd/apply"
	"github.com/proofzero/kmdr/cmd/setup"
	version "github.com/proofzero/kmdr/cmd/version"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kmdr",
	Short: "kmdr controls your kubelt applications.",
	Long:  `kmdr controls your kubelt applications.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO: init command to generate ~/.config/kubelt/config.toml
	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(apply.NewApplyCmd())
	rootCmd.AddCommand(setup.NewSetupCmd())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cf, _ := xdg.ConfigFile("kubelt/config.toml")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cf, fmt.Sprintf("config file (default is %s)", cf))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		_, err := xdg.ConfigFile("kubelt/config.toml")
		cobra.CheckErr(err)
		// Search config in ~/.config/kubelt directory with name "config" (without extension).
		viper.AddConfigPath(path.Join(xdg.ConfigHome, "kubelt"))

		// curDir, _ := os.Getwd()
		// viper.AddConfigPath(curDir)   // also load any config files in the current directory
		viper.SetConfigName("config") // home dir has kubelt yaml for secrets
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }
}
