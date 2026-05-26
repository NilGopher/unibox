package whoami

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	if len(args) > 0 {
		log.Print("usage: whoami")
		return 1
	}

	u, err := user.Current()
	if err != nil {
		log.Print(err)
		return 1
	}

	fmt.Fprintln(os.Stdout, u.Username)

	return 0
}
