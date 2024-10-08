package upgrade

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"github.com/nagylzs/gitlab-upgrade-artifact/internal/config"
	"io"
	"log/slog"
	"net/url"
	"os"
	"strconv"
)

type Upgrader struct {
	Opts     *config.CLIArgs
	artifact string // artifact path, in gitlab
	output   string // output path, local file name
	jobFile  string // job file path, a JSON that stores data about the local file's version
}

func (u *Upgrader) Upgrade() error {
	err := checkOpts(u.Opts)
	if err != nil {
		return err
	}
	u.artifact = u.Opts.Args.ArtifactName
	u.output = u.Opts.Args.OutputFile
	u.jobFile = u.Opts.Args.JsonFile
	if u.jobFile == "" {
		u.jobFile = u.output + ".job.json"
	}

	var programLevel = new(slog.LevelVar)
	if u.Opts.Debug {
		programLevel.Set(slog.LevelDebug)
	} else if u.Opts.Verbose {
		programLevel.Set(slog.LevelInfo)
	} else if u.Opts.Silent {
		programLevel.Set(slog.LevelError)
	} else {
		programLevel.Set(slog.LevelWarn)
	}
	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			NoColor: !isatty.IsTerminal(w.Fd()),
			Level:   programLevel,
		}),
	)
	slog.SetDefault(logger)

	// https://docs.gitlab.com/ee/api/jobs.html#list-project-jobs
	slug := url.PathEscape(u.Opts.Group + "/" + u.Opts.Project)
	jobListUrl := u.Opts.Server + "/api/v4/projects/" + slug + "/jobs"
	slog.Debug("Listing jobs", "url", jobListUrl)
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
	var artifactUrl string
	for idx, job := range jobs {
		artifactUrl = u.Opts.Server + "/api/v4/projects/" + slug + "/jobs/" + strconv.Itoa(job.Id) + "/artifacts/" + u.artifact
		slog.Debug("Examine job", "idx", idx, "url", artifactUrl)
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

	changed, err := u.fileChanged(commit)
	if err != nil {
		return err
	}
	if !changed && !u.Opts.Force {
		return nil
	}
	if u.Opts.Force {
		slog.Warn("Forcing artifact update")
	}

	slog.Debug("Creating local file", "path", u.output)
	out, err := os.Create(u.output)
	if err != nil {
		return fmt.Errorf("could not open output file '%v' for writing: %w", u.output, err)
	}
	defer out.Close()
	slog.Debug("Downloading local file", "path", u.output)
	err = getDownload(u, artifactUrl, out)
	if err != nil {
		return err
	}
	slog.Debug("Creating json file", "path", u.jobFile)
	jobFile, err := os.Create(u.jobFile)
	if err != nil {
		return fmt.Errorf("could not open output file '%v' for writing: %w", u.jobFile, err)
	}
	defer jobFile.Close()
	slog.Debug("Writing json file", "path", u.jobFile)
	data, err := json.Marshal(commit)
	if err != nil {
		return err
	}
	_, err = jobFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *Upgrader) fileChanged(commit JobListCommit) (bool, error) {
	jsonFile, err := os.Open(u.jobFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
		slog.Debug("No json file yet", "path", u.jobFile)
		return true, nil
	}
	defer jsonFile.Close()
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return false, fmt.Errorf("could not read data from %v: %w", u.jobFile, err)
	}
	var commitOld JobListCommit
	err = json.Unmarshal(data, &commitOld)
	if err != nil {
		slog.Warn("Could not unmarshal json file, assuming CHANGED", "path", u.jobFile, "err", err)
		return true, err
	}
	slog.Info("Old commit", "created_at", commit.CreatedAt, "hash", commitOld.Id, "message", commitOld.Message)
	slog.Info("New commit", "created_at", commit.CreatedAt, "hash", commit.Id, "message", commit.Message)
	changed := commitOld.Id != commit.Id
	if changed {
		slog.Info("CHANGED")
	} else {
		slog.Info("UNCHANGED")
	}
	return changed, nil
}
