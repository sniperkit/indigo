package cmd

import (
	"fmt"
	"github.com/comail/colog"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var logOutput *os.File

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the Indigo Server",
	Long:  `The start command starts the Indigo Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("start startCmd.PreRunE")

		fmt.Println(args)

		if len(args) < 1 {
			return cmd.Help()
		}

		_, _, err := cmd.Find(args)
		if err != nil {
			return err
		}

		fmt.Println("end startCmd.PersistentPreRunE")
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("start startCmd.RunE")

		fmt.Println(viper.GetString("output_format"))

		switch viper.GetString("output_format") {
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
			colog.SetFormatter(&colog.StdFormatter{
				Colors: false,
				Flag:   log.Ldate | log.Ltime | log.Lshortfile,
			})
		}

		if viper.GetString("log_output") != "" {
			var err error
			logOutput, err = os.OpenFile(viper.GetString("log_output"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			} else {
				colog.SetOutput(logOutput)
			}
		}

		switch viper.GetString("log_level") {
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

		fmt.Println("end startCmd.RunE")
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("start startCmd.PersistentPostRunE")

		if viper.GetString("log_output") != "" {
			logOutput.Close()
		}

		fmt.Println("end startCmd.PersistentPostRunE")
		return nil
	},
}

func init() {
	startCmd.PersistentFlags().StringVarP(&logOutputFile, "log-output", "o", constant.DefaultLogOutputFile, "log file")
	startCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", constant.DefaultLogLevel, "log level")

	viper.BindPFlag("log_output", startCmd.PersistentFlags().Lookup("log-output"))
	viper.BindPFlag("log_level", startCmd.PersistentFlags().Lookup("log-level"))

	RootCmd.AddCommand(startCmd)
}
