package cmd

import (
	"github.com/mosuka/indigo/rest"
	"github.com/mosuka/indigo/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		server := rest.NewIndigoRESTServer(IndigoSettings.GetInt("rest_port"), IndigoSettings.GetString("base_uri"), IndigoSettings.GetString("grpc_server"))
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
	startRESTCmd.Flags().IntP("port", "p", setting.DefaultRESTPort, "port number")
	startRESTCmd.Flags().StringP("base-uri", "b", setting.DefaultBaseURI, "base URI to run Indigo REST Server on")
	startRESTCmd.Flags().StringP("grpc-server", "g", setting.DefaultGRPCServer, "Indigo gRPC Sever")

	viper.BindPFlag("rest_port", startRESTCmd.Flags().Lookup("port"))
	viper.BindPFlag("base_uri", startRESTCmd.Flags().Lookup("base-uri"))
	viper.BindPFlag("grpc_server", startRESTCmd.Flags().Lookup("grpc-server"))

	startCmd.AddCommand(startRESTCmd)
}
