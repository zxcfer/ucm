package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/zxcfer/ucm/packages/commands"
)

func main() {
	if _, err := commands.Parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
