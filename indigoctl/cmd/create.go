package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
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
	CreateCmd.PersistentFlags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(CreateCmd)
}
