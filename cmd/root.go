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

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func loadConfig() {
	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		viper.SetConfigName("indigo")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("${HOME}/")
		viper.AddConfigPath(".")
	}
	viper.SetEnvPrefix("indigo")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func init() {
	cobra.OnInitialize(loadConfig)

	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", constant.DefaultConfigFile, "config file")
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", constant.DefaultOutputFormat, "output format")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", constant.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("output_format", RootCmd.PersistentFlags().Lookup("output-format"))
}
