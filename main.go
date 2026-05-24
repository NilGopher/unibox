package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/NilGopher/unibox/internal/cat"
	"github.com/NilGopher/unibox/internal/mkdir"
)

func main() {
	log.SetFlags(0)

	cmd, rawArgs := getArgs()

	switch cmd {
	case "cat":
		cat.Run(rawArgs)
	case "mkdir":
		mkdir.Run(rawArgs)
	default:
		log.Fatal("unknown command")
	}
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
