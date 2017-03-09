package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/grpc"
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
		server := grpc.NewIndigoGRPCServer(viper.GetInt("grpc_port"), viper.GetString("data_dir"))
		server.Start(viper.GetBool("open_existing_index"))

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
	startGRPCCmd.Flags().IntVarP(&gRPCPort, "port", "p", constant.DefaultGRPCPort, "port number")
	startGRPCCmd.Flags().StringVarP(&dataDir, "data-dir", "d", constant.DefaultDataDir, "data directory")
	startGRPCCmd.Flags().BoolVarP(&openExistingIndex, "open-existing-index", "e", constant.DefaultOpenExistingIndex, "open existing index")

	viper.BindPFlag("grpc_port", startGRPCCmd.Flags().Lookup("port"))
	viper.BindPFlag("data_dir", startGRPCCmd.Flags().Lookup("data-dir"))
	viper.BindPFlag("open_existing_index", startGRPCCmd.Flags().Lookup("open-existing-index"))

	startCmd.AddCommand(startGRPCCmd)
}
