package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/config"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/upgrade"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/version"
	"os"
)

func main() {
	var opts = config.CLIArgs{RequestTimeout: 10}
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	if opts.ShowVersion {
		version.PrintVersion()
		os.Exit(0)
	}

	err = upgrade.Upgrade(&opts, args)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

}
