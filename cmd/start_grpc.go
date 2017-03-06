package cmd

import (
	"github.com/mosuka/indigo/grpc"
	"github.com/mosuka/indigo/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		server := grpc.NewIndigoGRPCServer(IndigoSettings.GetInt("grpc_port"), IndigoSettings.GetString("data_dir"))
		server.Start(IndigoSettings.GetBool("open_existing_index"))

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
	startGRPCCmd.Flags().IntP("port", "p", setting.DefaultGRPCPort, "port number")
	startGRPCCmd.Flags().StringP("data-dir", "d", setting.DefaultDataDir, "data directory")
	startGRPCCmd.Flags().BoolP("open-existing-index", "O", setting.DefaultOpenExistingIndex, "open existing index")

	viper.BindPFlag("grpc_port", startGRPCCmd.Flags().Lookup("port"))
	viper.BindPFlag("data_dir", startGRPCCmd.Flags().Lookup("data-dir"))
	viper.BindPFlag("open_existing_index", startGRPCCmd.Flags().Lookup("open-existing-index"))

	startCmd.AddCommand(startGRPCCmd)
}
