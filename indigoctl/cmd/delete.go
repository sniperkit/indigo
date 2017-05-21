package cmd

import (
	"github.com/spf13/cobra"
)

type DeleteCommandOptions struct {
	gRPCServer string
}

var deleteCmdOpts DeleteCommandOptions

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes the object to the Indigo gRPC Server",
	Long:  `The delete command deletes the object to the Indigo gRPC Server.`,
	RunE:  runEDeleteCmd,
}

func runEDeleteCmd(cmd *cobra.Command, args []string) error {
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
	deleteCmd.PersistentFlags().StringVar(&deleteCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(deleteCmd)
}
