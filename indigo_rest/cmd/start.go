package cmd

import (
	"github.com/comail/colog"
	"github.com/mosuka/indigo/rest"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start Indigo REST Server",
	Long:  `The start command starts the Indigo REST Server.`,
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
		 * set server name
		 */
		serverName = config.GetString("server.name")

		/*
		 * set server port number
		 */
		serverPort = config.GetInt("server.port")

		/*
		 * set server path
		 */
		serverURIPath = config.GetString("server.uripath")

		/*
		 * set gRPC server port number
		 */
		gRPCServerName = config.GetString("grpc.server.name")

		/*
		 * set server port number
		 */
		gRPCServerPort = config.GetInt("grpc.server.port")

		/*
		 * start Indigo REST Server
		 */
		rs := rest.NewIndigoRESTServer(serverName, serverPort, serverURIPath, gRPCServerName, gRPCServerPort)
		rs.Start()

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
				rs.Stop()
				return nil
			case syscall.SIGINT:
				log.Println("info: trap SIGINT")
				rs.Stop()
				return nil
			case syscall.SIGTERM:
				log.Println("info: trap SIGTERM")
				rs.Stop()
				return nil
			case syscall.SIGQUIT:
				log.Println("info: trap SIGQUIT")
				rs.Stop()
				return nil
			default:
				log.Println("info: trap unknown signal")
				rs.Stop()
				return nil
			}
		}

		return nil
	},
}

func initConfig() {
	/*
	 * indigo_rest.yaml
	 */
	config.SetConfigName("indigo_rest")
	config.SetConfigType("yaml")
	config.AddConfigPath(configDir)
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("warn: %s\n", err.Error())
	}

	config.BindPFlag("log.file", startCmd.Flags().Lookup("log-file"))
	config.BindPFlag("log.level", startCmd.Flags().Lookup("log-level"))
	config.BindPFlag("log.format", startCmd.Flags().Lookup("log-format"))

	config.BindPFlag("server.name", startCmd.Flags().Lookup("server-name"))
	config.BindPFlag("server.port", startCmd.Flags().Lookup("server-port"))
	config.BindPFlag("server.uripath", startCmd.Flags().Lookup("server-uripath"))

	config.BindPFlag("grpc.server.name", startCmd.Flags().Lookup("grpc-server-name"))
	config.BindPFlag("grpc.server.port", startCmd.Flags().Lookup("grpc-server-port"))
}

func init() {
	cobra.OnInitialize(initConfig)

	startCmd.Flags().StringVarP(&configDir, "conf-dir", "c", configDir, "config directory")

	startCmd.Flags().StringVarP(&logFile, "log-file", "f", logFile, "log file")
	startCmd.Flags().StringVarP(&logLevel, "log-level", "l", logLevel, "log level")
	startCmd.Flags().StringVarP(&logFormat, "log-format", "F", logFormat, "log format")

	startCmd.Flags().StringVarP(&serverName, "server-name", "n", serverName, "name to run Indigo REST Server on")
	startCmd.Flags().IntVarP(&serverPort, "server-port", "p", serverPort, "port to run Indigo REST Server on")
	startCmd.Flags().StringVarP(&serverURIPath, "server-uripath", "u", serverURIPath, "URI path to run Indigo REST Server on")

	startCmd.Flags().StringVarP(&gRPCServerName, "grpc-server-name", "N", gRPCServerName, "name to run Indigo gRPC Server on")
	startCmd.Flags().IntVarP(&gRPCServerPort, "grpc-server-port", "P", gRPCServerPort, "port to run Indigo gRPC Server on")

	RootCmd.AddCommand(startCmd)
}
