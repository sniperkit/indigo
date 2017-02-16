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
		gs := grpc.NewIndigoGRPCServer(serverPort, dataDir)
		gs.Start()

		/*
		 * trap signals
		 */
		signal_chan := make(chan os.Signal, 1)
		signal.Notify(signal_chan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				log.Println("info: trap SIGHUP")
				gs.Stop()
				return nil
			case syscall.SIGINT:
				log.Println("info: trap SIGINT")
				gs.Stop()
				return nil
			case syscall.SIGTERM:
				log.Println("info: trap SIGTERM")
				gs.Stop()
				return nil
			case syscall.SIGQUIT:
				log.Println("info: trap SIGQUIT")
				gs.Stop()
				return nil
			default:
				log.Println("info: trap unknown")
				gs.Stop()
				return nil
			}
		}

		return nil
	},
}

func init() {
	startGRPCCmd.Flags().IntVarP(&serverPort, "server-port", "p", serverPort, "port number")
	startGRPCCmd.Flags().StringVarP(&dataDir, "data-dir", "d", dataDir, "data directory")

	startCmd.AddCommand(startGRPCCmd)
}
