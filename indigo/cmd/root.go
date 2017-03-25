package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	ver "github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var RootCmd = &cobra.Command{
	Use:               "indigo",
	Short:             "CLI for Indigo Server",
	Long:              `The Command Line Interface for the Indigo gRPC or REST Server.`,
	PersistentPreRunE: persistentPreRunERootCmd,
	RunE:              runERootCmd,
}

func persistentPreRunERootCmd(cmd *cobra.Command, args []string) error {
	if versionFlag {
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
	viper.SetDefault("log_output_format", constant.DefaultLogOutputFormat)
	viper.SetDefault("log_output", constant.DefaultLogOutput)
	viper.SetDefault("log_level", constant.DefaultLogLevel)

	viper.SetDefault("grpc.port", constant.DefaultGRPCPort)
	viper.SetDefault("grpc.data_dir", constant.DefaultDataDir)
	viper.SetDefault("grpc.open_existing_index", constant.DefaultOpenExistingIndex)

	viper.SetDefault("rest.port", constant.DefaultRESTPort)
	viper.SetDefault("rest.base_uri", constant.DefaultBaseURI)
	viper.SetDefault("rest.grpc_server", constant.DefaultGRPCServer)

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

	RootCmd.PersistentFlags().StringP("config", "c", constant.DefaultConfig, "configuration file of Indigo Server")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", constant.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
