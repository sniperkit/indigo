package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "puts the object from the Indigo gRPC Server",
	Long:  `The put command puts the object from the Indigo gRPC Server.`,
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
	putCmd.PersistentFlags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")

	viper.BindPFlag("grpc_server", putCmd.PersistentFlags().Lookup("grpc-server"))

	RootCmd.AddCommand(putCmd)
}
