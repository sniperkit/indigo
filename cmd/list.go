package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the object from the Indigo gRPC Server",
	Long:  `The list command lists the object from the Indigo gRPC Server.`,
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
	listCmd.PersistentFlags().StringP("grpc-server", "g", indigoSettings.GetString("grpc_server"), "Indigo gRPC Sever")
	indigoSettings.BindPFlag("grpc_server", listCmd.Flags().Lookup("grpc-server"))

	RootCmd.AddCommand(listCmd)
}
