package touch

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	if len(args) == 0 {
		log.Print("usage: touch <options> <file_name>...")
		return 1
	}

	mode, err := strconv.ParseUint(opt.mode, 8, 32)
	if err != nil {
		log.Print(err)
		return 1
	}

	for _, arg := range args {
		file, err := os.OpenFile(arg, os.O_CREATE|os.O_EXCL, 0)
		if err != nil {
			if os.IsExist(err) {
				os.Chtimes(arg, time.Now(), time.Now())
			} else {
				log.Print(err)
				exitCode = 1
			}
			continue
		}

		file.Chmod(os.FileMode(mode))
		file.Close()
	}

	return
}
