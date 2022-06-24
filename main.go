package main

import (
	"github.com/dihedron/seal/command"
	"github.com/jessevdk/go-flags"

	"os"
)

func main() {
	cmd := command.Commands{}
	if _, err := flags.NewParser(&cmd, flags.Default).Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			//fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		default:
			//fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

}
