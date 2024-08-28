package upgrade

import (
	"fmt"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/config"
	"os"
	"strings"
)

func checkOpts(opts *config.CLIArgs) error {
	if opts.Server == "" {
		opts.Server = os.Getenv("GITLAB_SERVER")
		if opts.Server == "" {
			return fmt.Errorf("--server is required")
		}
		if !strings.HasPrefix(opts.Server, "https://") && strings.HasPrefix(opts.Server, "http://") {
			return fmt.Errorf("--server must start with https:// or http://")
		}
	}
	if opts.Token == "" {
		opts.Token = os.Getenv("GITLAB_TOKEN")
		if opts.Token == "" {
			return fmt.Errorf("--token is required")
		}
	}
	if opts.Group == "" {
		opts.Group = os.Getenv("GITLAB_GROUP")
		if opts.Group == "" {
			return fmt.Errorf("--group is required")
		}
	}
	if opts.Project == "" {
		opts.Project = os.Getenv("GITLAB_PROJECT")
		if opts.Project == "" {
			return fmt.Errorf("--project is required")
		}
	}
	if opts.Branch == "" {
		opts.Branch = os.Getenv("GITLAB_BRANCH")
		if opts.Branch == "" {
			opts.Branch = "main"
		}
	}
	if opts.Job == "" {
		opts.Job = os.Getenv("GITLAB_JOB")
		if opts.Job == "" {
			return fmt.Errorf("--job is required")
		}
	}
	if opts.Debug && opts.Verbose {
		return fmt.Errorf("cannot combine --debug and --verbose")
	}

	if opts.RequestTimeout < 1 {
		return fmt.Errorf("--request-timeout cannot be less than 1 (second)")
	}

	if opts.DownloadTimeout < 5 {
		return fmt.Errorf("--download-timeout cannot be less than 5 (second)")
	}

	return nil
}
