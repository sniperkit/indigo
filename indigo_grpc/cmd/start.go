package cmd

import (
	"fmt"
	"github.com/comail/colog"
	g "github.com/mosuka/indigo/grpc"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start Indigo gRPC Server",
	Long:  `The start command starts the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * set log file
		 */
		logFile = config.GetString("log.file")
		f, err := os.OpenFile(
			logFile,
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
		logLevel = config.GetString("log.level")
		switch logLevel {
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
		switch logFormat {
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
		 * set grpc port number
		 */
		serverName = config.GetString("server.name")

		/*
		 * set grpc port number
		 */
		serverPort = config.GetInt("server.port")

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
		 * start Bleve Server asynchronously
		 */
		server := grpc.NewServer()
		proto.RegisterIndigoServer(server, g.NewIndigoGRPCServer(indexDir, indexMapping, indexType, indexStore))
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return err
		}

		go func() {
			log.Printf("info: start server name=%s port=%d\n", serverName, serverPort)
			server.Serve(listener)
			return
		}()

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
				log.Printf("info: stop server name=%s port=%d trap=SIGHUP\n", serverName, serverPort)
				server.GracefulStop()
				return nil
			case syscall.SIGINT:
				log.Printf("info: stop server name=%s port=%d trap=SIGINT\n", serverName, serverPort)
				server.GracefulStop()
				return nil
			case syscall.SIGTERM:
				log.Printf("info: stop server name=%s port=%d trap=SIGTERM\n", serverName, serverPort)
				server.GracefulStop()
				return nil
			case syscall.SIGQUIT:
				log.Printf("info: stop server name=%s port=%d trap=SIGQUIT\n", serverName, serverPort)
				server.GracefulStop()
				return nil
			default:
				log.Printf("info: stop server name=%s port=%d trap=unknown\n", serverName, serverPort)
				server.GracefulStop()
				return nil
			}
		}

		return nil
	},
}

func initConfig() {
	/*
	 * indigo_grpc.yaml
	 */
	config.SetConfigName("indigo_grpc")
	config.SetConfigType("yaml")
	config.AddConfigPath(configDir)
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("warn: %s\n", err.Error())
	}

	config.BindPFlag("log.file", startCmd.Flags().Lookup("log-file"))
	config.BindPFlag("log.level", startCmd.Flags().Lookup("log-level"))
	config.BindPFlag("log.format", startCmd.Flags().Lookup("log-format"))

	config.BindPFlag("grpc.name", startCmd.Flags().Lookup("grpc-name"))
	config.BindPFlag("grpc.port", startCmd.Flags().Lookup("grpc-port"))

	config.BindPFlag("index.dir", startCmd.Flags().Lookup("index-dir"))
	config.BindPFlag("index.type", startCmd.Flags().Lookup("index-type"))
	config.BindPFlag("index.store", startCmd.Flags().Lookup("index-store"))
	config.BindPFlag("index.mapping", startCmd.Flags().Lookup("index-mapping"))
}

func init() {
	cobra.OnInitialize(initConfig)

	startCmd.Flags().StringVarP(&configDir, "conf-dir", "c", configDir, "config directory")

	startCmd.Flags().StringVarP(&logFile, "log-file", "f", logFile, "log file")
	startCmd.Flags().StringVarP(&logLevel, "log-level", "l", logLevel, "log level")
	startCmd.Flags().StringVarP(&logFormat, "log-format", "F", logFormat, "log format")

	startCmd.Flags().StringVarP(&serverName, "server-name", "n", serverName, "name to run Indigo Server on")
	startCmd.Flags().IntVarP(&serverPort, "server-port", "p", serverPort, "port to run Indigo Server on")

	startCmd.Flags().StringVarP(&indexDir, "index-dir", "d", indexDir, "index path")
	startCmd.Flags().StringVarP(&indexMapping, "index-mapping", "m", indexMapping, "index mapping")
	startCmd.Flags().StringVarP(&indexType, "index-type", "t", indexType, "index type")
	startCmd.Flags().StringVarP(&indexStore, "index-store", "s", indexStore, "index store")

	RootCmd.AddCommand(startCmd)
}
