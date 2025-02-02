package flag

import "flag"

type NewsLetterFlag struct {
	ConfigPath string
}

func ParseFlags(arguments []string) (NewsLetterFlag, bool, error) {
	flagSet := flag.NewFlagSet("newsletter-generator", flag.ExitOnError)
	flags := NewsLetterFlag{}
	flagSet.StringVar(
		&flags.ConfigPath, "config-file", "config/dev.yaml", "config file path",
	)
	dryRun := flagSet.Bool(
		"dry-run", false, "for dry-run",
	)
	if err := flagSet.Parse(arguments); err != nil {
		return NewsLetterFlag{}, false, err
	}

	return flags, *dryRun, nil
}
