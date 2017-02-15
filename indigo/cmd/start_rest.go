package cmd

import (
	"github.com/mosuka/indigo/rest"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var startRESTCmd = &cobra.Command{
	Use:   "rest",
	Short: "start Indigo REST Server",
	Long:  `The rest command starts the Indigo REST Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * start Indigo REST Server
		 */
		rs := rest.NewIndigoRESTServer(serverPort, restServerURIPath, grpcServerName, grpcServerPort)
		rs.Start()

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
				rs.Stop()
				return nil
			case syscall.SIGINT:
				log.Println("info: trap SIGINT")
				rs.Stop()
				return nil
			case syscall.SIGTERM:
				log.Println("info: trap SIGTERM")
				rs.Stop()
				return nil
			case syscall.SIGQUIT:
				log.Println("info: trap SIGQUIT")
				rs.Stop()
				return nil
			default:
				log.Println("info: trap unknown signal")
				rs.Stop()
				return nil
			}
		}

		return nil
	},
}

func init() {
	startRESTCmd.Flags().StringVarP(&restServerURIPath, "server-uripath", "u", restServerURIPath, "URI path to run Indigo REST Server on")

	startRESTCmd.Flags().StringVarP(&grpcServerName, "grpc-server-name", "N", grpcServerName, "name to run Indigo gRPC Server on")
	startRESTCmd.Flags().IntVarP(&grpcServerPort, "grpc-server-port", "P", grpcServerPort, "port to run Indigo gRPC Server on")

	startCmd.AddCommand(startRESTCmd)
}
