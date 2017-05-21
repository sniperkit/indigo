package cmd

import (
	"github.com/spf13/cobra"
)

type PutCommandOptions struct {
	gRPCServer string
}

var putCmdOpts PutCommandOptions

var putCmd = &cobra.Command{
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
	putCmd.PersistentFlags().StringVar(&putCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(putCmd)
}
