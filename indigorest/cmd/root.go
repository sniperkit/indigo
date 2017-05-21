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
	Use:               "indigorest",
	Short:             "CLI for Indigo REST Server",
	Long:              `The Command Line Interface for the Indigo REST Server.`,
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
	viper.SetDefault("base_uri", DefaultBaseURI)
	viper.SetDefault("server", DefaultServer)

	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		viper.SetConfigName("indigorest")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/indigo")
		viper.AddConfigPath("${HOME}/indigo")
		viper.AddConfigPath("./indigo")
	}
	viper.SetEnvPrefix("indigorest")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func init() {
	cobra.OnInitialize(LoadConfig)

	RootCmd.PersistentFlags().String("config", DefaultConfig, "configuration file of Indigo Server")
	RootCmd.PersistentFlags().BoolVar(&rootCmdOpts.versionFlag, "version", false, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
