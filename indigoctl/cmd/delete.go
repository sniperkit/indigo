package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
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
	DeleteCmd.PersistentFlags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(DeleteCmd)
}
