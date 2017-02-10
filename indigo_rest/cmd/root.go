package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "indigo_rest",
	Short: "Indigo REST Server",
	Long:  `The Indigo REST Server is a gateway server which provides REST API for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("must specify command or flags")
		}
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
