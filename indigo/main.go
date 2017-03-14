package main

import (
	"github.com/mosuka/indigo/indigo/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
