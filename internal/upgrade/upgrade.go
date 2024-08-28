package upgrade

import (
	"errors"
	"fmt"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/config"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Upgrader struct {
	Opts     *config.CLIArgs
	Args     []string
	artifact string // artifact path, in gitlab
	output   string // output path, local file name
	jobFile  string // job file path, a JSON that stores data about the local file's version
	cl       *http.Client
}

func (u *Upgrader) Upgrade() error {
	err := checkOpts(u.Opts)
	if err != nil {
		return err
	}
	if len(u.Args) < 3 || len(u.Args) > 4 {
		return errors.New("positional arguments are: artifact_path output_path [job_file]")
	}
	u.artifact = u.Args[1]
	u.output = u.Args[2]
	if len(u.Args) > 3 {
		u.jobFile = u.Args[3]
	} else {
		u.jobFile = u.output + ".job.json"
	}

	u.cl = &http.Client{Timeout: time.Second * time.Duration(u.Opts.RequestTimeout)}

	// https://docs.gitlab.com/ee/api/jobs.html#list-project-jobs
	slug := url.PathEscape(u.Opts.Group + "/" + u.Opts.Project)
	jobListUrl := u.Opts.Server + "/api/v4/projects/" + slug + "/jobs"
	var jobs []JobListItem
	err = getAndDecode(u, jobListUrl, &jobs)
	if err != nil {
		return err
	}

	// We go over the jobs here, and get the artifact from the latest one.
	// https://docs.gitlab.com/ee/api/job_artifacts.html#download-a-single-artifact-file-from-specific-tag-or-branch
	// GET /projects/:id/jobs/artifacts/:ref_name/raw/*artifact_path?job=name
	// https://docs.gitlab.com/ee/api/job_artifacts.html#download-a-single-artifact-file-by-job-id
	// GET /projects/:id/jobs/:job_id/artifacts/*artifact_path
	var commit JobListCommit
	ok := false
	for _, job := range jobs {
		artifactUrl := u.Opts.Server + "/api/v4/projects/" + slug + "/jobs/" + strconv.Itoa(job.Id) + "/artifacts/" + u.artifact
		r, err := head(u, artifactUrl)
		if err != nil {
			return err
		}
		if r.StatusCode == 404 {
			continue
		}
		if r.StatusCode != 200 {
			return errors.New(r.Status)
		}
		commit = job.Commit
		ok = true
		break
	}
	if !ok {
		return errors.New("artifact '" + u.artifact + "' not found in any job")
	}
	fmt.Println(commit)
	// TODO: compare/fetch/upgrade here

	return nil
}
