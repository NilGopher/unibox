package cat

import "flag"

type options struct {
	number bool
}

var opt = &options{}

func (o *options) parse(rawArgs []string) (args []string) {
	flags := flag.NewFlagSet("cat", flag.ExitOnError)

	flags.BoolVar(&o.number, "n", false, "number output lines")
	flags.BoolVar(&o.number, "number", false, "number output lines")

	flags.Parse(rawArgs)

	return flags.Args()
}
