package mkdir

import (
	"log"
	"os"
	"strconv"
)

func Run(rawArgs []string) {
	args := opt.parse(rawArgs)

	if len(args) == 0 {
		log.Fatal("usage: mkdir <options> <dir_name>...")
	}

	mode, err := strconv.ParseInt(opt.mode, 8, 32)
	if err != nil {
		log.Fatal(err)
	}

	for _, arg := range args {
		if opt.parents {
			os.MkdirAll(arg, os.FileMode(mode))
		} else {
			os.MkdirAll(arg, os.FileMode(mode))
		}
	}
}
