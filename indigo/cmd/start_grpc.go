package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/indigo/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

var StartGRPCCmd = &cobra.Command{
	Use:   "grpc",
	Short: "starts Indigo gRPC Server",
	Long:  `The start grpc command starts the Indigo gRPC Server.`,
	RunE:  runEStartGRPCCmd,
}

func runEStartGRPCCmd(cmd *cobra.Command, args []string) error {
	server := grpc.NewIndigoGRPCServer(viper.GetInt("grpc.port"), viper.GetString("grpc.data_dir"))
	server.Start(viper.GetBool("grpc.open_existing_index"))

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
	StartGRPCCmd.Flags().IntP("port", "p", constant.DefaultGRPCPort, "port number to be used when Indigo gRPC Server starts up")
	StartGRPCCmd.Flags().StringP("data-dir", "d", constant.DefaultDataDir, "path of the directory where Indigo gRPC Server stores the data")
	StartGRPCCmd.Flags().BoolP("open-existing-index", "e", constant.DefaultOpenExistingIndex, "flag to open indices when started to Indigo gRPC Server")

	viper.BindPFlag("grpc.port", StartGRPCCmd.Flags().Lookup("port"))
	viper.BindPFlag("grpc.data_dir", StartGRPCCmd.Flags().Lookup("data-dir"))
	viper.BindPFlag("grpc.open_existing_index", StartGRPCCmd.Flags().Lookup("open-existing-index"))

	StartCmd.AddCommand(StartGRPCCmd)
}
