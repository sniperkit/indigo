package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/indigo/grpc"
	"github.com/mosuka/indigo/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
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

	// IndexMapping
	indexMapping := mapping.NewIndexMapping()
	if cmd.Flag("index-mapping").Changed {
		file, err := os.Open(viper.GetString("index_mapping"))
		if err != nil {
			return err
		}
		defer file.Close()

		resourceBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(resourceBytes, indexMapping)
		if err != nil {
			return err
		}
	}

	// Kvconfig
	kvconfig := make(map[string]interface{})
	if cmd.Flag("kvconfig").Changed {
		file, err := os.Open(viper.GetString("kvconfig"))
		if err != nil {
			return err
		}
		defer file.Close()

		resourceBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(resourceBytes, kvconfig)
		if err != nil {
			return err
		}
	}

	server := grpc.NewIndigoGRPCServer(viper.GetInt("port"), viper.GetString("path"), indexMapping, viper.GetString("index_type"), viper.GetString("kvstore"), kvconfig)
	server.Start(viper.GetBool("delete_index_at_startup"))

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

		server.Stop(viper.GetBool("delete_index_at_shutdown"))

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
	StartCmd.Flags().String("log-format", DefaultLogFormat, "log format")
	StartCmd.Flags().String("log-output", DefaultLogOutput, "log output path")
	StartCmd.Flags().String("log-level", DefaultLogLevel, "log level")
	StartCmd.Flags().Int("port", DefaultPort, "port number")
	StartCmd.Flags().String("path", DefaultPath, "index directory path")
	StartCmd.Flags().String("index-mapping", DefaultIndexMapping, "index mapping path")
	StartCmd.Flags().String("index-type", DefaultIndexType, "index type")
	StartCmd.Flags().String("kvstore", DefaultKvstore, "kvstore")
	StartCmd.Flags().String("kvconfig", DefaultKvconfig, "kvconfig path")
	StartCmd.Flags().Bool("delete-index-at-startup", DefaultDeleteIndexAtStartup, "delete index at startup")
	StartCmd.Flags().Bool("delete-index-at-shutdown", DefaultDeleteIndexAtShutdown, "delete index at shutdown")

	viper.BindPFlag("log_format", StartCmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log_output", StartCmd.Flags().Lookup("log-output"))
	viper.BindPFlag("log_level", StartCmd.Flags().Lookup("log-level"))
	viper.BindPFlag("port", StartCmd.Flags().Lookup("port"))
	viper.BindPFlag("path", StartCmd.Flags().Lookup("path"))
	viper.BindPFlag("index_mapping", StartCmd.Flags().Lookup("index-mapping"))
	viper.BindPFlag("index_type", StartCmd.Flags().Lookup("index-type"))
	viper.BindPFlag("kvstore", StartCmd.Flags().Lookup("kvstore"))
	viper.BindPFlag("kvconfig", StartCmd.Flags().Lookup("kvconfig"))
	viper.BindPFlag("delete_index_at_startup", StartCmd.Flags().Lookup("delete-index-at-startup"))
	viper.BindPFlag("delete_index_at_shutdown", StartCmd.Flags().Lookup("delete-index-at-shutdown"))

	RootCmd.AddCommand(StartCmd)
}
