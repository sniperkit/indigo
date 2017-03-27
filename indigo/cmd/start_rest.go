package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/indigo/rest"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

var StartRESTCmd = &cobra.Command{
	Use:   "rest",
	Short: "start Indigo REST Server",
	Long:  `The start rest command starts the Indigo REST Server.`,
	RunE:  runEStartRESTCmd,
}

func runEStartRESTCmd(cmd *cobra.Command, args []string) error {
	server := rest.NewIndigoRESTServer(viper.GetInt("rest.port"), viper.GetString("rest.base_uri"), viper.GetString("rest.grpc_server"))
	server.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	for {
		sig := <-signalChan

		log.WithFields(log.Fields{
			"signal": sig,
		}).Info("trap signal")

		server.Stop()

		return nil
	}

	return nil
}

func init() {
	StartRESTCmd.Flags().IntP("port", "p", constant.DefaultRESTPort, "port number to be used when Indigo REST Server starts up")
	StartRESTCmd.Flags().StringP("base-uri", "b", constant.DefaultBaseURI, "base URI of API endpoint on Indigo REST Server")
	StartRESTCmd.Flags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC server that Indigo REST Server connect to")

	viper.BindPFlag("rest.port", StartRESTCmd.Flags().Lookup("port"))
	viper.BindPFlag("rest.base_uri", StartRESTCmd.Flags().Lookup("base-uri"))
	viper.BindPFlag("rest.grpc_server", StartRESTCmd.Flags().Lookup("grpc-server"))

	StartCmd.AddCommand(StartRESTCmd)
}
