package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/defaultvalue"
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
	viper.SetDefault("log_format", defaultvalue.DefaultLogFormat)
	viper.SetDefault("log_output", defaultvalue.DefaultLogOutput)
	viper.SetDefault("log_level", defaultvalue.DefaultLogLevel)

	viper.SetDefault("grpc.port", defaultvalue.DefaultGRPCPort)
	viper.SetDefault("grpc.data_dir", defaultvalue.DefaultDataDir)
	viper.SetDefault("grpc.open_existing_index", defaultvalue.DefaultOpenExistingIndex)

	viper.SetDefault("rest.port", defaultvalue.DefaultRESTPort)
	viper.SetDefault("rest.base_uri", defaultvalue.DefaultBaseURI)
	viper.SetDefault("rest.grpc_server", defaultvalue.DefaultGRPCServer)

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

	RootCmd.PersistentFlags().String("config", defaultvalue.DefaultConfig, "configuration file of Indigo Server")
	RootCmd.PersistentFlags().BoolVar(&versionFlag, "version", defaultvalue.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
