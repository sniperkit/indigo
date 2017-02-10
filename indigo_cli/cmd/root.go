package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "indigo_cli",
	Short: "Indigo Command Line Interface",
	Long:  `The Indigo Command Line Interface is controlls the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
