package touch

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Run(rawArgs []string) {
	args := opt.parse(rawArgs)

	if len(args) == 0 {
		log.Fatal("usage: touch <options> <file_name>...")
	}

	mode, err := strconv.ParseUint(opt.mode, 8, 32)
	if err != nil {
		log.Fatal(err)
	}

	for _, arg := range args {
		file, err := os.OpenFile(arg, os.O_CREATE|os.O_EXCL, 0)
		if err != nil {
			if os.IsExist(err) {
				os.Chtimes(arg, time.Now(), time.Now())
			} else {
				log.Print(err)
			}
			continue
		}

		file.Chmod(os.FileMode(mode))
		file.Close()
	}
}
