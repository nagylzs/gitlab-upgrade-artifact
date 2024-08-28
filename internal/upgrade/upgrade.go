package upgrade

import (
	"encoding/json"
	"fmt"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/config"
	"net/http"
	"net/url"
	"time"
)

func Upgrade(opts *config.CLIArgs, args []string) error {
	err := checkOpts(opts)
	if err != nil {
		return err
	}

	cl := &http.Client{Timeout: time.Second * time.Duration(opts.RequestTimeout)}

	// https://docs.gitlab.com/ee/api/jobs.html#list-project-jobs
	slug := url.PathEscape(opts.Group + "/" + opts.Project)
	jobListUrl := opts.Server + "/api/v4/projects/" + slug + "/jobs"

	req, err := http.NewRequest("GET", jobListUrl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("PRIVATE-TOKEN", opts.Token)
	r, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var jobs []JobListItem
	err = json.NewDecoder(r.Body).Decode(&jobs)
	if err != nil {
		return err
	}
	fmt.Println(jobs)

	return nil
}
