package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	ver "github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:               "indigo",
	Short:             "CLI for controling Indigo Server",
	Long:              `The Command Line Interface for controling the Indigo Server.`,
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
	RootCmd.PersistentFlags().StringVar(&outputFormat, "output-format", constant.DefaultOutputFormat, "output format of the command execution result")
	RootCmd.PersistentFlags().BoolVar(&versionFlag, "verson", constant.DefaultVersionFlag, "show version numner")
}
