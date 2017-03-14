package main

import (
	"github.com/mosuka/indigo/indigoctl/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
