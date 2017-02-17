package cmd

import (
	"errors"
	"github.com/comail/colog"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var logOutput *os.File

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the Indigo Server",
	Long:  `The start command starts the Indigo Server.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		/*
		 * set log file
		 */
		if logOutputFile != "" {
			logOutput, err := os.OpenFile(logOutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			} else {
				colog.SetOutput(logOutput)
			}
		}

		/*
		 * set log level
		 */
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
		default:
			colog.SetFormatter(&colog.StdFormatter{
				Colors: true,
				Flag:   log.Ldate | log.Ltime | log.Lshortfile,
			})
		}

		colog.ParseFields(true)

		colog.Register()

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if logOutputFile != "" {
			logOutput.Close()
		}

		return nil
	},
}

func init() {
	startCmd.PersistentFlags().StringVarP(&logOutputFile, "log-output-file", "o", logOutputFile, "log file")
	startCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", logLevel, "log level")
	startCmd.PersistentFlags().StringVarP(&logFormat, "log-format", "f", logFormat, "log format")

	RootCmd.AddCommand(startCmd)
}
