package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "indigo_grpc",
	Short: "Indigo gRPC Server",
	Long:  `The Indigo gRPC Server is a search server built on top of the Bleve.`,
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
