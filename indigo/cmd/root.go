//  Copyright (c) 2015 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	ver "github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type RootCommandOptions struct {
	versionFlag bool
}

var rootCmdOpts RootCommandOptions

var RootCmd = &cobra.Command{
	Use:               "indigo",
	Short:             "CLI for Indigo Server",
	Long:              `The Command Line Interface for the Indigo Server.`,
	PersistentPreRunE: persistentPreRunERootCmd,
	RunE:              runERootCmd,
}

func persistentPreRunERootCmd(cmd *cobra.Command, args []string) error {
	if rootCmdOpts.versionFlag {
		fmt.Printf("%s\n", ver.Version)
		os.Exit(0)
	}

	return nil
}

func runERootCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Help()
	}

	return nil
}

func LoadConfig() {
	viper.SetDefault("log_format", DefaultLogFormat)
	viper.SetDefault("log_output", DefaultLogOutput)
	viper.SetDefault("log_level", DefaultLogLevel)

	viper.SetDefault("port", DefaultPort)
	viper.SetDefault("path", DefaultPath)
	viper.SetDefault("index_mapping", DefaultIndexMapping)
	viper.SetDefault("index_type", DefaultIndexType)
	viper.SetDefault("kvstore", DefaultKvstore)
	viper.SetDefault("kvconfig", DefaultKvconfig)
	viper.SetDefault("open_existing_index", DefaultDeleteIndexAtStartup)
	viper.SetDefault("delete_index_at_shutdown", DefaultDeleteIndexAtShutdown)

	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		viper.SetConfigName("indigo")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/indigo")
		viper.AddConfigPath("${HOME}/indigo")
		viper.AddConfigPath("./indigo")
	}
	viper.SetEnvPrefix("indigo")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func init() {
	cobra.OnInitialize(LoadConfig)

	RootCmd.PersistentFlags().String("config", DefaultConfig, "configuration file path")
	RootCmd.PersistentFlags().BoolVar(&rootCmdOpts.versionFlag, "version", false, "show version number")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
