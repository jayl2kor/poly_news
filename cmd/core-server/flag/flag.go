package flag

import "flag"

type CoreServerFlag struct {
	ConfigPath string
}

func ParseFlags(arguments []string) (CoreServerFlag, error) {
	flagSet := flag.NewFlagSet("core-server", flag.ExitOnError)
	flags := CoreServerFlag{}
	flagSet.StringVar(
		&flags.ConfigPath, "config-file", "config/dev.yaml", "config file path",
	)
	if err := flagSet.Parse(arguments); err != nil {
		return CoreServerFlag{}, err
	}

	return flags, nil
}
