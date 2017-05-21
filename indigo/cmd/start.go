package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/indigo/grpc"
	"github.com/mosuka/indigo/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logOutput *os.File

var StartCmd = &cobra.Command{
	Use:                "start",
	Short:              "starts the Indigo Server",
	Long:               `The start command starts the Indigo Server.`,
	PersistentPreRunE:  persistentPreRunEStartCmd,
	RunE:               runEStartCmd,
	PersistentPostRunE: persistentPostRunEStartCmd,
}

func persistentPreRunEStartCmd(cmd *cobra.Command, args []string) error {
	if rootCmdOpts.versionFlag {
		fmt.Printf("%s\n", version.Version)
		os.Exit(0)
	}

	switch viper.GetString("log_format") {
	case "text":
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      false,
			DisableColors:    true,
			DisableTimestamp: false,
			FullTimestamp:    true,
			TimestampFormat:  time.RFC3339,
			DisableSorting:   false,
			QuoteEmptyFields: true,
			QuoteCharacter:   "\"",
		})
	case "color":
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      true,
			DisableColors:    false,
			DisableTimestamp: false,
			FullTimestamp:    true,
			TimestampFormat:  time.RFC3339,
			DisableSorting:   false,
			QuoteEmptyFields: true,
			QuoteCharacter:   "\"",
		})
	case "json":
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat:  time.RFC3339,
			DisableTimestamp: false,
			FieldMap: log.FieldMap{
				log.FieldKeyTime:  "@timestamp",
				log.FieldKeyLevel: "@level",
				log.FieldKeyMsg:   "@message",
			},
		})
	default:
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      false,
			DisableColors:    true,
			DisableTimestamp: false,
			FullTimestamp:    true,
			TimestampFormat:  time.RFC3339,
			DisableSorting:   false,
			QuoteEmptyFields: true,
			QuoteCharacter:   "\"",
		})
	}

	switch viper.GetString("log_level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	if viper.GetString("log_output") == "" {
		log.SetOutput(os.Stdout)
	} else {
		var err error
		logOutput, err = os.OpenFile(viper.GetString("log_output"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		} else {
			log.SetOutput(logOutput)
		}
	}

	return nil
}

func runEStartCmd(cmd *cobra.Command, args []string) error {
	server := grpc.NewIndigoGRPCServer(viper.GetInt("port"), viper.GetString("data_dir"))
	server.Start(viper.GetBool("open_existing_index"))

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

func persistentPostRunEStartCmd(cmd *cobra.Command, args []string) error {
	if viper.GetString("log_output") != "" {
		logOutput.Close()
	}

	return nil
}

func init() {
	StartCmd.Flags().String("log-format", DefaultLogFormat, "log format of Indigo Server")
	StartCmd.Flags().String("log-output", DefaultLogOutput, "log output destination of Indigo Server")
	StartCmd.Flags().String("log-level", DefaultLogLevel, "log level of log output by Indigo Server")
	StartCmd.Flags().Int("port", DefaultPort, "port number to be used when Indigo gRPC Server starts up")
	StartCmd.Flags().String("data-dir", DefaultDataDir, "path of the directory where Indigo gRPC Server stores the data")
	StartCmd.Flags().Bool("open-existing-index", DefaultOpenExistingIndex, "flag to open indices when started to Indigo gRPC Server")

	viper.BindPFlag("log_format", StartCmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log_output", StartCmd.Flags().Lookup("log-output"))
	viper.BindPFlag("log_level", StartCmd.Flags().Lookup("log-level"))
	viper.BindPFlag("port", StartCmd.Flags().Lookup("port"))
	viper.BindPFlag("data_dir", StartCmd.Flags().Lookup("data-dir"))
	viper.BindPFlag("open_existing_index", StartCmd.Flags().Lookup("open-existing-index"))

	RootCmd.AddCommand(StartCmd)
}
