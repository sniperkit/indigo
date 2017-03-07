package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/setting"
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

	fmt.Printf("config: %s\n", viper.GetString("config"))

	fmt.Printf("config: %s\n", configFile)
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

	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", setting.DefaultConfigFile, "config file")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", setting.DefaultOutputFormat, "output format")
	viper.BindPFlag("output_format", RootCmd.PersistentFlags().Lookup("output-format"))

	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "show version numner")
}
