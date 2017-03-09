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
		fmt.Println("start RootCmd.PersistentPreRunE")

		if versionFlag {
			fmt.Printf("%s\n", version.Version)
			os.Exit(0)
		}

		fmt.Println("end RootCmd.PersistentPreRunE")
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("start RootCmd.RunE")

		if len(args) < 1 {
			return cmd.Help()
		}

		fmt.Println("end RootCmd.RunE")
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func loadConfig() {
	fmt.Println("start loadConfig")

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

	fmt.Println("end loadConfig")
}

func init() {
	cobra.OnInitialize(loadConfig)

	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", constant.DefaultConfigFile, "config file")
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", constant.DefaultOutputFormat, "output format")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", constant.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("output_format", RootCmd.PersistentFlags().Lookup("output-format"))

	fmt.Println("end RootCmd init")
}
