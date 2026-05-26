package pwd

import (
	"fmt"
	"log"
	"os"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	if len(args) > 0 {
		log.Print("usage: pwd")
		return 1
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return 1
	}

	fmt.Fprintln(os.Stdout, wd)

	return
}
