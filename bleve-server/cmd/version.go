package cmd

import (
	"fmt"
	"github.com/mosuka/bleve-server/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version number",
	Long:  `The version command shows Bleve Server's version number.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s\n", version.Version)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
