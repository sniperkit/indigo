package cmd

import (
	"github.com/mosuka/indigo/constant"
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
	Long:  `The start rest command starts the Indigo REST Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * start Indigo REST Server
		 */
		server := rest.NewIndigoRESTServer(restServerPort, baseURI, gRPCServer)
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
				log.Println("info: trap unknown signal")
				server.Stop()
				return nil
			}
		}

		return nil
	},
}

func init() {
	startRESTCmd.Flags().IntVarP(&restServerPort, "server-port", "p", constant.DefaultRESTServerPort, "port number")
	startRESTCmd.Flags().StringVarP(&baseURI, "base-uri", "b", constant.DefaultBaseURI, "base URI to run Indigo REST Server on")
	startRESTCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")

	startCmd.AddCommand(startRESTCmd)
}
