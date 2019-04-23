package main

import (
	"os"

	"github.com/mrubelmann/bb8beat/cmd"

	_ "github.com/mrubelmann/bb8beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
