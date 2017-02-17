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
		server := grpc.NewIndigoGRPCServer(gRPCServerPort, dataDir)
		server.Start()

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
	startGRPCCmd.Flags().IntVarP(&gRPCServerPort, "server-port", "p", gRPCServerPort, "port number")
	startGRPCCmd.Flags().StringVarP(&dataDir, "data-dir", "d", dataDir, "data directory")

	startCmd.AddCommand(startGRPCCmd)
}
