package cmd

import (
	"github.com/spf13/cobra"
)

type CloseCommandOptions struct {
	gRPCServer string
}

var closeCmdOpts CloseCommandOptions

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "closes the object to the Indigo gRPC Server",
	Long:  `The open command creates the object to the Indigo gRPC Server.`,
	RunE:  runECloseCmd,
}

func runECloseCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Help()
	}

	_, _, err := cmd.Find(args)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	closeCmd.PersistentFlags().StringVar(&closeCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(closeCmd)
}
