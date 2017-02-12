package cmd

import (
	"github.com/comail/colog"
	"github.com/mosuka/indigo/grpc"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "start Indigo gRPC Server",
	Long:  `The grpc command starts the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * set log file
		 */
		grpcLogFile = config.GetString("grpc.log.file")
		f, err := os.OpenFile(
			grpcLogFile,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0644)
		if err != nil {
			log.Println(err.Error())
		} else {
			colog.SetOutput(f)
		}
		defer f.Close()

		/*
		 * set log level
		 */
		grpcLogLevel = config.GetString("grpc.log.level")
		switch grpcLogLevel {
		case "trace":
			colog.SetDefaultLevel(colog.LTrace)
		case "debug":
			colog.SetDefaultLevel(colog.LDebug)
		case "info":
			colog.SetDefaultLevel(colog.LInfo)
		case "warn":
			colog.SetDefaultLevel(colog.LWarning)
		case "error":
			colog.SetDefaultLevel(colog.LError)
		case "alert":
			colog.SetDefaultLevel(colog.LAlert)
		default:
			colog.SetDefaultLevel(colog.LInfo)
		}

		/*
		 * set log format
		 */
		grpcLogLevel = config.GetString("grpc.log.format")
		switch grpcLogFormat {
		case "text":
			colog.SetFormatter(&colog.StdFormatter{
				Colors: true,
				Flag:   log.Ldate | log.Ltime | log.Lshortfile,
			})
		case "json":
			colog.SetFormatter(&colog.JSONFormatter{
				TimeFormat: time.RFC3339,
				Flag:       log.Lshortfile,
			})
		}

		colog.Register()

		/*
		 * set server port number
		 */
		grpcServerName = config.GetString("grpc.server.name")

		/*
		 * set server port number
		 */
		grpcServerPort = config.GetInt("grpc.server.port")

		/*
		 * set index directory
		 */
		indexDir = config.GetString("index.dir")

		/*
		 * set index mapping
		 */
		indexMapping = config.GetString("index.mapping")

		/*
		 * set index type
		 */
		indexType = config.GetString("index.type")

		/*
		 * set index store
		 */
		indexStore = config.GetString("index.store")

		/*
		 * start Indigo gRPC Server
		 */
		gs := grpc.NewIndigoGRPCServer(grpcServerName, grpcServerPort, indexDir, indexMapping, indexType, indexStore)
		gs.Start()

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
				gs.Stop()
				return nil
			case syscall.SIGINT:
				log.Println("info: trap SIGINT")
				gs.Stop()
				return nil
			case syscall.SIGTERM:
				log.Println("info: trap SIGTERM")
				gs.Stop()
				return nil
			case syscall.SIGQUIT:
				log.Println("info: trap SIGQUIT")
				gs.Stop()
				return nil
			default:
				log.Println("info: trap unknown")
				gs.Stop()
				return nil
			}
		}

		return nil
	},
}

func initGRPCCmd() {
	/*
	 * indigo.yaml
	 */
	config.SetConfigName("indigo")
	config.SetConfigType("yaml")
	config.AddConfigPath(configDir)
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("warn: %s\n", err.Error())
	}

	config.BindPFlag("grpc.server.name", grpcCmd.Flags().Lookup("grpc-name"))
	config.BindPFlag("grpc.server.port", grpcCmd.Flags().Lookup("grpc-port"))
	config.BindPFlag("grpc.log.file", grpcCmd.Flags().Lookup("log-file"))
	config.BindPFlag("grpc.log.level", grpcCmd.Flags().Lookup("log-level"))
	config.BindPFlag("grpc.log.format", grpcCmd.Flags().Lookup("log-format"))

	config.BindPFlag("index.dir", grpcCmd.Flags().Lookup("index-dir"))
	config.BindPFlag("index.type", grpcCmd.Flags().Lookup("index-type"))
	config.BindPFlag("index.store", grpcCmd.Flags().Lookup("index-store"))
	config.BindPFlag("index.mapping", grpcCmd.Flags().Lookup("index-mapping"))
}

func init() {
	cobra.OnInitialize(initGRPCCmd)

	grpcCmd.Flags().StringVarP(&configDir, "conf-dir", "c", configDir, "grpcConfig directory")

	grpcCmd.Flags().StringVarP(&grpcServerName, "server-name", "n", grpcServerName, "name to run Indigo gRPC Server on")
	grpcCmd.Flags().IntVarP(&grpcServerPort, "server-port", "p", grpcServerPort, "port to run Indigo gRPC Server on")
	grpcCmd.Flags().StringVarP(&grpcLogFile, "log-file", "f", grpcLogFile, "log file")
	grpcCmd.Flags().StringVarP(&grpcLogLevel, "log-level", "l", grpcLogLevel, "log level")
	grpcCmd.Flags().StringVarP(&grpcLogFormat, "log-format", "F", grpcLogFormat, "log format")

	grpcCmd.Flags().StringVarP(&indexDir, "index-dir", "d", indexDir, "index path")
	grpcCmd.Flags().StringVarP(&indexMapping, "index-mapping", "m", indexMapping, "index mapping")
	grpcCmd.Flags().StringVarP(&indexType, "index-type", "t", indexType, "index type")
	grpcCmd.Flags().StringVarP(&indexStore, "index-store", "s", indexStore, "index store")

	RootCmd.AddCommand(grpcCmd)
}
