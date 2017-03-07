package cmd

import (
	"github.com/mosuka/indigo/rest"
	"github.com/mosuka/indigo/constant"
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
		server := rest.NewIndigoRESTServer(viper.GetInt("rest_port"), viper.GetString("base_uri"), viper.GetString("grpc_server"))
		server.Start()

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
	startRESTCmd.Flags().IntVarP(&restPort, "port", "p", constant.DefaultRESTPort, "port number")
	viper.BindPFlag("rest_port", startRESTCmd.Flags().Lookup("port"))

	startRESTCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	viper.BindPFlag("grpc_server", startRESTCmd.Flags().Lookup("grpc-server"))

	startRESTCmd.Flags().StringVarP(&baseURI, "base-uri", "b", constant.DefaultBaseURI, "base URI to run Indigo REST Server on")
	viper.BindPFlag("base_uri", startRESTCmd.Flags().Lookup("base-uri"))

	startCmd.AddCommand(startRESTCmd)
}
