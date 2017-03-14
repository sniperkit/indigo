package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "indigo",
	Short: "Indigo Command Line Interface",
	Long:  `The Indigo Command Line Interface controlls the Indigo Server.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("%s\n", version.Version)
			os.Exit(0)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		return nil
	},
}

func LoadConfig() {
	viper.SetDefault("output_format", constant.DefaultOutputFormat)
	viper.SetDefault("log_output", constant.DefaultLogOutput)
	viper.SetDefault("log_level", constant.DefaultLogLevel)
	viper.SetDefault("grpc_port", constant.DefaultGRPCPort)
	viper.SetDefault("data_dir", constant.DefaultDataDir)
	viper.SetDefault("open_existing_index", constant.DefaultOpenExistingIndex)
	viper.SetDefault("batch_size", constant.DefaultBatchSize)
	viper.SetDefault("index", constant.DefaultIndex)
	viper.SetDefault("rest_port", constant.DefaultRESTPort)
	viper.SetDefault("grpc_server", constant.DefaultGRPCServer)
	viper.SetDefault("base_uri", constant.DefaultBaseURI)
	viper.SetDefault("index", constant.DefaultIndex)

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

	RootCmd.PersistentFlags().StringP("config", "c", constant.DefaultConfig, "config file")
	RootCmd.PersistentFlags().StringP("output-format", "f", constant.DefaultOutputFormat, "output format")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", constant.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("output_format", RootCmd.PersistentFlags().Lookup("output-format"))
}
