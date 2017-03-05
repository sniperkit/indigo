package cmd

import (
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "closes the object to the Indigo gRPC Server",
	Long:  `The open command creates the object to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		_, _, err := cmd.Find(args)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	closeCmd.PersistentFlags().StringP("grpc-server", "g", indigoSettings.GetString("grpc_server"), "Indigo gRPC Sever")
	indigoSettings.BindPFlag("grpc_server", closeCmd.Flags().Lookup("grpc-server"))

	RootCmd.AddCommand(closeCmd)
}
