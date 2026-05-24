package touch

import "flag"

type options struct {
	mode string
}

var opt = &options{}

func (o *options) parse(rawArgs []string) (args []string) {
	flags := flag.NewFlagSet("touch", flag.ExitOnError)

	flags.StringVar(&o.mode, "mode", "644", "set file mode")

	flags.Parse(rawArgs)

	return flags.Args()
}
