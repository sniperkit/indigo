package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the object from the Indigo gRPC Server",
	Long:  `The get command gets the object from the Indigo gRPC Server.`,
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
	getCmd.PersistentFlags().StringP("grpc-server", "g", indigoSettings.GetString("grpc_server"), "Indigo gRPC Sever")
	indigoSettings.BindPFlag("grpc_server", getCmd.Flags().Lookup("grpc-server"))

	RootCmd.AddCommand(getCmd)
}
