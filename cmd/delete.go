package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes the object to the Indigo gRPC Server",
	Long:  `The delete command deletes the object to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		_, _, err := cmd.Find(args)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	deleteCmd.PersistentFlags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	viper.BindPFlag("grpc_server", deleteCmd.PersistentFlags().Lookup("grpc-server"))

	RootCmd.AddCommand(deleteCmd)
}
