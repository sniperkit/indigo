package cmd

import (
	"github.com/spf13/cobra"
)

type CreateCommandOptions struct {
	gRPCServer string
}

var createCmdOpts CreateCommandOptions

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates the object to the Indigo gRPC Server",
	Long:  `The create command creates the object to the Indigo gRPC Server.`,
	RunE:  runECreateCmd,
}

func runECreateCmd(cmd *cobra.Command, args []string) error {
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
	createCmd.PersistentFlags().StringVar(&createCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(createCmd)
}
