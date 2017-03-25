package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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
	if versionFlag {
		fmt.Printf("%s\n", version.Version)
		os.Exit(0)
	}

	switch viper.GetString("log_output_format") {
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
	if len(args) < 1 {
		return cmd.Help()
	}

	_, _, err := cmd.Find(args)
	if err != nil {
		return err
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
	StartCmd.PersistentFlags().StringP("log-output-format", "f", constant.DefaultLogOutputFormat, "log output format of Indigo Server")
	StartCmd.PersistentFlags().StringP("log-output", "o", constant.DefaultLogOutput, "log output destination of Indigo Server")
	StartCmd.PersistentFlags().StringP("log-level", "l", constant.DefaultLogLevel, "log level of log output by Indigo Server")

	viper.BindPFlag("log_output_format", StartCmd.PersistentFlags().Lookup("log-output-format"))
	viper.BindPFlag("log_output", StartCmd.PersistentFlags().Lookup("log-output"))
	viper.BindPFlag("log_level", StartCmd.PersistentFlags().Lookup("log-level"))

	RootCmd.AddCommand(StartCmd)
}
