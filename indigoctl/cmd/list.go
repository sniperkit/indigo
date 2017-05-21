package cmd

import (
	"github.com/spf13/cobra"
)

type ListCommandOptions struct {
	gRPCServer string
}

var listCmdOpts ListCommandOptions

var listCmd = &cobra.Command{
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
	listCmd.PersistentFlags().StringVar(&listCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(listCmd)
}
