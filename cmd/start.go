package cmd

import (
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
		//outputFormat = indigoSettings.GetString("output_format")
		switch indigoSettings.GetString("output_format") {
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

		/*
		 * set log file
		 */
		if indigoSettings.GetString("log_output") != "" {
			logOutput, err := os.OpenFile(indigoSettings.GetString("log_output"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			} else {
				colog.SetOutput(logOutput)
			}
		}

		/*
		 * set log level
		 */
		switch indigoSettings.GetString("log_level") {
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
		if indigoSettings.GetString("log_output") != "" {
			logOutput.Close()
		}

		return nil
	},
}

func init() {
	startCmd.PersistentFlags().StringP("log-output", "o", indigoSettings.GetString("log_output"), "log file")
	indigoSettings.BindPFlag("log_output", RootCmd.Flags().Lookup("log-output"))

	startCmd.PersistentFlags().StringP("log-level", "l", indigoSettings.GetString("log_level"), "log level")
	indigoSettings.BindPFlag("log_level", RootCmd.Flags().Lookup("log-level"))

	RootCmd.AddCommand(startCmd)
}
