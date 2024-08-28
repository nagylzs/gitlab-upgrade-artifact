
# gitlab-upgrade-artifact

This little tool can compare the local version and the remote version of a gitlab CI/CD generated artifact,
and download a new version if it is available.

It creates a json file, containing commit information of the downloaded artifact. It goes over the CI/CD jobs of 
the given artifact, finds the last commit that resulted in (re)building the artifact, and compares the remote
version's commit hash with the local one. The name of the json file is  output_file_name + ".job.json" by default,
but it can be changed by giving the third argument.

## Usage

```
Usage:
  gitlab-upgrade-artifact [OPTIONS] artifact_name output_file_name [json_file_name]

Application Options:
  -s, --server=           Gitlab server, defaults to GITLAB_SERVER env
  -t, --token=            Gitlab token, defaults to GITLAB_TOKEN env
  -g, --group=            Gitlab project group, defaults to GITLAB_GROUP env
  -p, --project=          Gitlab project, defaults to GITLAB_PROJECT env
  -b, --branch=           Gitlab branch, defaults to GITLAB_BRANCH env, or
                          'main'
  -j, --job=              Gitlab job, defaults to GITLAB_JOB env
  -f, --force             Force download, even if not changed
      --silent            Be silent, only errors are logged
  -v, --verbose           Be verbose, warnings and errors are logged
  -d, --debug             Debug messages
      --version           Show version information and exit
      --request-timeout=  Request timeout in seconds, for JSON API http
                          requests (default: 10)
      --download-timeout= Download timeout in seconds, for artifact download
                          http requests (default: 120)

Help Options:
  -h, --help              Show this help message

Usage:
  gitlab-upgrade-artifact [OPTIONS]

Application Options:
  -s, --server=           Gitlab server, defaults to GITLAB_SERVER env
  -t, --token=            Gitlab token, defaults to GITLAB_TOKEN env
  -g, --group=            Gitlab project group, defaults to GITLAB_GROUP env
  -p, --project=          Gitlab project, defaults to GITLAB_PROJECT env
  -b, --branch=           Gitlab branch, defaults to GITLAB_BRANCH env, or
                          'main'
  -j, --job=              Gitlab job, defaults to GITLAB_JOB env
  -f, --force             Force download, even if not changed
      --silent            Be silent, only errors are logged
  -v, --verbose           Be verbose, warnings and errors are logged
  -d, --debug             Debug messages
      --version           Show version information and exit
      --request-timeout=  Request timeout in seconds, for JSON API http
                          requests (default: 10)
      --download-timeout= Download timeout in seconds, for artifact download
                          http requests (default: 120)

Help Options:
  -h, --help              Show this help message
```

## Example output

```bash
gitlab-upgrade-artifact --debug --server https://gitlab.example.com --group mygroup --token fflat-2CsEvyb5zJacWcevHhhs --project my-project-name --branch dev --job build_all --force dist/linux/amd64/example-artifact-name example-artifact-name
Aug 28 18:41:38.419 DBG Listing jobs url=https://gitlab.example.com/api/v4/projects/mygroup%2Fmy-project-name/jobs
Aug 28 18:41:39.324 DBG Examine job idx=0 url=https://gitlab.example.com/api/v4/projects/mygroup%2Fmy-project-name/jobs/5360/artifacts/dist/linux/amd64/example-artifact-name
Aug 28 18:41:39.884 INF Old commit created_at="2024-08-27 18:52:46 +0200 CEST" hash=7ee7720f5c90640fd480bcc4ddb41aae2616d383 message="start dev branch\n"
Aug 28 18:41:39.884 INF New commit created_at="2024-08-27 18:52:46 +0200 CEST" hash=7ee7720f5c90640fd480bcc4ddb41aae2616d383 message="start dev branch\n"
Aug 28 18:41:39.884 INF UNCHANGED
Aug 28 18:41:39.884 WRN Forcing artifact update
Aug 28 18:41:39.884 DBG Creating local file path=example-artifact-name
Aug 28 18:41:39.885 DBG Downloading local file path=example-artifact-name
downloading 100% |███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| (20/20 MB, 5.4 MB/s)        
Aug 28 18:41:43.851 DBG Creating json file path=example-artifact-name.job.json
Aug 28 18:41:43.851 DBG Writing json file path=example-artifact-name.job.json

Process finished with the exit code 0
```