package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/NilGopher/unibox/internal/cat"
	"github.com/NilGopher/unibox/internal/mkdir"
	"github.com/NilGopher/unibox/internal/rm"
	"github.com/NilGopher/unibox/internal/touch"
)

func main() {
	log.SetFlags(0)

	var exitCode int

	cmd, rawArgs := getArgs()

	switch cmd {
	case "cat":
		exitCode = cat.Run(rawArgs)
	case "mkdir":
		exitCode = mkdir.Run(rawArgs)
	case "rm":
		exitCode = rm.Run(rawArgs)
	case "touch":
		exitCode = touch.Run(rawArgs)
	default:
		log.Print("unknown command:", cmd)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func getArgs() (cmd string, rawArgs []string) {
	cmd = filepath.Base(os.Args[0])

	if len(os.Args) > 1 {
		if cmd == "unibox" {
			cmd = os.Args[1]
			rawArgs = os.Args[2:]
		} else {
			rawArgs = os.Args[1:]
		}
	}

	return
}
