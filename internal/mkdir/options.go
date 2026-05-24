package mkdir

import "flag"

type options struct {
	mode    string
	parents bool
}

var opt = &options{}

func (o *options) parse(rawArgs []string) (args []string) {
	flags := flag.NewFlagSet("mkdir", flag.ExitOnError)

	flags.StringVar(&o.mode, "m", "755", "set file mode")
	flags.StringVar(&o.mode, "mode", "755", "set file mode")

	flags.BoolVar(&o.parents, "p", false, "make parent directories as needed")
	flags.BoolVar(&o.parents, "parents", false, "make parent directories as needed")

	flags.Parse(rawArgs)

	return flags.Args()
}
