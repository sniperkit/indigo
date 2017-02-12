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

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "start Indigo REST Server",
	Long:  `The rest command starts the Indigo REST Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * set log file
		 */
		restLogFile = config.GetString("rest.log.file")
		f, err := os.OpenFile(
			restLogFile,
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
		restLogLevel = config.GetString("rest.log.level")
		switch restLogLevel {
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
		restLogFormat = config.GetString("rest.log.format")
		switch restLogFormat {
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
		restServerName = config.GetString("rest.server.name")

		/*
		 * set server port number
		 */
		restServerPort = config.GetInt("rest.server.port")

		/*
		 * set server path
		 */
		restServerURIPath = config.GetString("rest.server.uripath")

		/*
		 * set gRPC server port number
		 */
		grpcServerName = config.GetString("grpc.server.name")

		/*
		 * set server port number
		 */
		grpcServerPort = config.GetInt("grpc.server.port")

		/*
		 * start Indigo REST Server
		 */
		rs := rest.NewIndigoRESTServer(restServerName, restServerPort, restServerURIPath, grpcServerName, grpcServerPort)
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
	config.SetConfigName("indigo")
	config.SetConfigType("yaml")
	config.AddConfigPath(configDir)
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("warn: %s\n", err.Error())
	}

	config.BindPFlag("rest.log.file", restCmd.Flags().Lookup("log-file"))
	config.BindPFlag("rest.log.level", restCmd.Flags().Lookup("log-level"))
	config.BindPFlag("rest.log.format", restCmd.Flags().Lookup("log-format"))

	config.BindPFlag("rest.server.name", restCmd.Flags().Lookup("server-name"))
	config.BindPFlag("rest.server.port", restCmd.Flags().Lookup("server-port"))
	config.BindPFlag("rest.server.uripath", restCmd.Flags().Lookup("server-uripath"))

	config.BindPFlag("grpc.server.name", restCmd.Flags().Lookup("grpc-server-name"))
	config.BindPFlag("grpc.server.port", restCmd.Flags().Lookup("grpc-server-port"))
}

func init() {
	cobra.OnInitialize(initConfig)

	restCmd.Flags().StringVarP(&configDir, "conf-dir", "c", configDir, "config directory")

	restCmd.Flags().StringVarP(&restLogFile, "log-file", "f", restLogFile, "log file")
	restCmd.Flags().StringVarP(&restLogLevel, "log-level", "l", restLogLevel, "log level")
	restCmd.Flags().StringVarP(&restLogFormat, "log-format", "F", restLogFormat, "log format")

	restCmd.Flags().StringVarP(&restServerName, "server-name", "n", restServerName, "name to run Indigo REST Server on")
	restCmd.Flags().IntVarP(&restServerPort, "server-port", "p", restServerPort, "port to run Indigo REST Server on")
	restCmd.Flags().StringVarP(&restServerURIPath, "server-uripath", "u", restServerURIPath, "URI path to run Indigo REST Server on")

	restCmd.Flags().StringVarP(&grpcServerName, "grpc-server-name", "N", grpcServerName, "name to run Indigo gRPC Server on")
	restCmd.Flags().IntVarP(&grpcServerPort, "grpc-server-port", "P", grpcServerPort, "port to run Indigo gRPC Server on")

	RootCmd.AddCommand(restCmd)
}
