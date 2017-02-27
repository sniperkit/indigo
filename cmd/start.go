package cmd

import (
	"github.com/comail/colog"
	"github.com/mosuka/indigo/constant"
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
			colog.SetMinLevel(colog.LTrace)
		case "debug":
			colog.SetMinLevel(colog.LDebug)
		case "info":
			colog.SetMinLevel(colog.LInfo)
		case "warn":
			colog.SetMinLevel(colog.LWarning)
		case "error":
			colog.SetMinLevel(colog.LError)
		case "alert":
			colog.SetMinLevel(colog.LAlert)
		default:
			colog.SetMinLevel(colog.LInfo)
		}

		/*
		 * set log format
		 */
		switch outputFormat {
		case "text":
			colog.SetFormatter(&colog.StdFormatter{
				Colors: false,
				Flag:   log.Ldate | log.Ltime | log.Lshortfile,
			})
		case "color":
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
			// same as text
			colog.SetFormatter(&colog.StdFormatter{
				Colors: false,
				Flag:   log.Ldate | log.Ltime | log.Lshortfile,
			})
		}

		colog.SetDefaultLevel(colog.LInfo)

		colog.ParseFields(true)

		colog.Register()

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		_, _, err := cmd.Find(args)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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
	startCmd.PersistentFlags().StringVarP(&logOutputFile, "log-output", "o", constant.DefaultLogOutputFile, "log file")
	startCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", constant.DefaultLogLevel, "log level")

	RootCmd.AddCommand(startCmd)
}
