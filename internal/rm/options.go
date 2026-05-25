package rm

import "flag"

type options struct {
	force bool
}

var opt = &options{}

func (o *options) parse(rawArgs []string) (args []string) {
	flags := flag.NewFlagSet("rm", flag.ExitOnError)

	flags.BoolVar(&o.force, "f", false, "force deletion")
	flags.BoolVar(&o.force, "force", false, "force deletion")

	flags.Parse(rawArgs)

	return flags.Args()
}
