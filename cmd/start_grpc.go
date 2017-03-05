package cmd

import (
	"github.com/mosuka/indigo/grpc"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var startGRPCCmd = &cobra.Command{
	Use:   "grpc",
	Short: "start Indigo gRPC Server",
	Long:  `The start grpc command starts the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * start Indigo gRPC Server
		 */
		server := grpc.NewIndigoGRPCServer(indigoSettings.GetInt("grpc_port"), indigoSettings.GetString("data_dir"))
		server.Start(indigoSettings.GetBool("open_existing_index"))

		/*
		 * trap signals
		 */
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		for {
			sig := <-signalChan
			switch sig {
			case syscall.SIGHUP:
				log.Println("info: trap SIGHUP")
				server.Stop()
				return nil
			case syscall.SIGINT:
				log.Println("info: trap SIGINT")
				server.Stop()
				return nil
			case syscall.SIGTERM:
				log.Println("info: trap SIGTERM")
				server.Stop()
				return nil
			case syscall.SIGQUIT:
				log.Println("info: trap SIGQUIT")
				server.Stop()
				return nil
			default:
				log.Println("info: trap unknown")
				server.Stop()
				return nil
			}
		}

		return nil
	},
}

func init() {
	startGRPCCmd.Flags().IntP("port", "p", indigoSettings.GetInt("grpc_port"), "port number")
	indigoSettings.BindPFlag("grpc_port", startGRPCCmd.Flags().Lookup("port"))

	startGRPCCmd.Flags().StringP("data-dir", "d", indigoSettings.GetString("data_dir"), "data directory")
	indigoSettings.BindPFlag("data_dir", startGRPCCmd.Flags().Lookup("data-dir"))

	startGRPCCmd.Flags().BoolP("open-existing-index", "O", indigoSettings.GetBool("open_existing_index"), "open existing index")
	indigoSettings.BindPFlag("open_existing_index", startGRPCCmd.Flags().Lookup("open-existing-index"))

	startCmd.AddCommand(startGRPCCmd)
}
