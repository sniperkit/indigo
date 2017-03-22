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
	Short:             "Indigo Command Line Interface",
	Long:              `The Indigo Command Line Interface controlls the Indigo Server.`,
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

func init() {
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", constant.DefaultOutputFormat, "output format")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "persistentPreRunERootCmd", "v", constant.DefaultVersionFlag, "show version numner")

	viper.BindPFlag("output_format", RootCmd.PersistentFlags().Lookup("output-format"))
}
