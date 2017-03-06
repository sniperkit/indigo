package cmd

import (
	"fmt"
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
		fmt.Println("startGRPCCmd.RunE")

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
	startGRPCCmd.Flags().IntVarP(&gRPCPort, "port", "p", setting.DefaultGRPCPort, "port number")
	viper.BindPFlag("grpc_port", startGRPCCmd.Flags().Lookup("port"))

	startGRPCCmd.Flags().StringVarP(&dataDir, "data-dir", "d", setting.DefaultDataDir, "data directory")
	viper.BindPFlag("data_dir", startGRPCCmd.Flags().Lookup("data-dir"))

	startGRPCCmd.Flags().BoolVarP(&openExistingIndex, "open-existing-index", "e", setting.DefaultOpenExistingIndex, "open existing index")
	viper.BindPFlag("open_existing_index", startGRPCCmd.Flags().Lookup("open-existing-index"))

	startCmd.AddCommand(startGRPCCmd)
}
