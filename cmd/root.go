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
		fmt.Println("RootCmd.PersistentPreRunE")

		if versionFlag {
			fmt.Printf("%s\n", version.Version)
			os.Exit(0)
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("RootCmd.PreRunE")

		if len(args) < 1 {
			return cmd.Help()
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("RootCmd.RunE")

		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func loadConfig() {
	if configFile == "" {
		viper.SetConfigName("indigo")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("${HOME}/")
		viper.AddConfigPath("./")
	} else {
		viper.SetConfigFile(configFile)
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func init() {
	cobra.OnInitialize(loadConfig)

	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", setting.DefaultConfigFile, "config file")

	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", setting.DefaultOutputFormat, "output format")
	viper.BindPFlag("output_format", RootCmd.Flags().Lookup("output-format"))

	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "show version numner")
}
