package cmd

import (
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/spf13/cobra"
)

var PutCmd = &cobra.Command{
	Use:   "put",
	Short: "puts the object from the Indigo gRPC Server",
	Long:  `The put command puts the object from tlhe Indigo gRPC Server.`,
	RunE:  runEPutCmd,
}

func runEPutCmd(cmd *cobra.Command, args []string) error {
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
	PutCmd.PersistentFlags().StringVar(&gRPCServer, "grpc-server", defaultvalue.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(PutCmd)
}
