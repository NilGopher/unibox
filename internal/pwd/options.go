package pwd

import "flag"

type options struct {
}

var opt = &options{}

func (o *options) parse(rawArgs []string) (args []string) {
	flags := flag.NewFlagSet("pwd", flag.ExitOnError)

	flags.Parse(rawArgs)

	return flags.Args()
}
