package cmd

import (
	"github.com/spf13/cobra"
)

type GetCommandOptions struct {
	gRPCServer string
}

var getCmdOpts GetCommandOptions

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the object from the Indigo gRPC Server",
	Long:  `The get command gets the object from the Indigo gRPC Server.`,
	RunE:  runEGetCmd,
}

func runEGetCmd(cmd *cobra.Command, args []string) error {
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
	getCmd.PersistentFlags().StringVar(&getCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(getCmd)
}
