package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the object from the Indigo gRPC Server",
	Long:  `The get command gets the object from the Indigo gRPC Server.`,
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
	getCmd.PersistentFlags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	viper.BindPFlag("grpc_server", getCmd.PersistentFlags().Lookup("grpc-server"))

	RootCmd.AddCommand(getCmd)
}
