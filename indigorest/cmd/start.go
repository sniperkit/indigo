//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/indigorest/rest"
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
	Short:              "starts the Indigo REST Server",
	Long:               `The start command starts the Indigo REST Server.`,
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
	server := rest.NewIndigoRESTServer(viper.GetInt("port"), viper.GetString("base_uri"), viper.GetString("server"))
	server.Start()

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
	StartCmd.Flags().Int("port", DefaultPort, "port number to be used when Indigo REST Server starts up")
	StartCmd.Flags().String("base-uri", DefaultBaseURI, "base URI of API endpoint on Indigo REST Server")
	StartCmd.Flags().String("server", DefaultServer, "Indigo gRPC server that Indigo REST Server connect to")

	viper.BindPFlag("log_format", StartCmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log_output", StartCmd.Flags().Lookup("log-output"))
	viper.BindPFlag("log_level", StartCmd.Flags().Lookup("log-level"))
	viper.BindPFlag("port", StartCmd.Flags().Lookup("port"))
	viper.BindPFlag("base_uri", StartCmd.Flags().Lookup("base-uri"))
	viper.BindPFlag("server", StartCmd.Flags().Lookup("server"))

	RootCmd.AddCommand(StartCmd)
}
