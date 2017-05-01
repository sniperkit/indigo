package cmd

import (
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the object from the Indigo gRPC Server",
	Long:  `The list command lists the object from the Indigo gRPC Server.`,
	RunE:  runEListCmd,
}

func runEListCmd(cmd *cobra.Command, args []string) error {
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
	ListCmd.PersistentFlags().StringVar(&gRPCServer, "grpc-server", defaultvalue.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(ListCmd)
}
