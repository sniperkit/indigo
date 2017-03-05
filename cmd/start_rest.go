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
	Long:  `The start rest command starts the Indigo REST Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * start Indigo REST Server
		 */
		server := rest.NewIndigoRESTServer(indigoSettings.GetInt("rest_port"), indigoSettings.GetString("base_uri"), indigoSettings.GetString("grpc_server"))
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
	startRESTCmd.Flags().IntP("port", "p", indigoSettings.GetInt("rest_port"), "port number")
	indigoSettings.BindPFlag("rest_port", startRESTCmd.Flags().Lookup("port"))

	startRESTCmd.Flags().StringP("base-uri", "b", indigoSettings.GetString("base_uri"), "base URI to run Indigo REST Server on")
	indigoSettings.BindPFlag("base_uri", startRESTCmd.Flags().Lookup("base-uri"))

	startRESTCmd.Flags().StringP("grpc-server", "g", indigoSettings.GetString("grpc_server"), "Indigo gRPC Sever")
	indigoSettings.BindPFlag("grpc_server", startRESTCmd.Flags().Lookup("grpc-server"))

	startRESTCmd.Flags().StringP("grpc-server", "g", indigoSettings.GetString("grpc_server"), "Indigo gRPC Sever")
	indigoSettings.BindPFlag("grpc_server", startRESTCmd.Flags().Lookup("grpc-server"))

	startCmd.AddCommand(startRESTCmd)
}
