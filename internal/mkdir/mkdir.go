package mkdir

import (
	"log"
	"os"
	"strconv"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	if len(args) == 0 {
		log.Print("usage: mkdir <options> <dir_name>...")
		return 1
	}

	mode, err := strconv.ParseUint(opt.mode, 8, 32)
	if err != nil {
		log.Print(err)
		return 1
	}

	for _, arg := range args {
		if opt.parents {
			if err := os.MkdirAll(arg, 0); err != nil {
				log.Print(err)
				exitCode = 1
				continue
			}
		} else {
			if err := os.Mkdir(arg, 0); err != nil {
				log.Print(err)
				exitCode = 1
				continue
			}
		}

		if err := os.Chmod(arg, os.FileMode(mode)); err != nil {
			log.Print(err)
			exitCode = 1
		}
	}

	return
}
