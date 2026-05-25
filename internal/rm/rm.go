package rm

import (
	"log"
	"os"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	if len(args) == 0 {
		log.Print("usage: rm <options> <targets>")
		return 1
	}

	for _, arg := range args {
		if opt.force {
			if err := os.RemoveAll(arg); err != nil {
				log.Print(err)
				exitCode = 1
			}
		} else {
			if err := os.Remove(arg); err != nil {
				log.Print(err)
				exitCode = 1
			}
		}
	}

	return
}
