package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/version"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "indigo",
	Short: "Indigo Command Line Interface",
	Long:  `The Indigo Command Line Interface controlls the Indigo Server.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("%s\n", version.Ver)
			os.Exit(0)
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", constant.DefaultVersionFlag, "show version numner")
}
