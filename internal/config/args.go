package config

type CLIArgs struct {
	Server          string `short:"s" long:"server" description:"Gitlab server, defaults to GITLAB_SERVER env"`
	Token           string `short:"t" long:"token" description:"Gitlab token, defaults to GITLAB_TOKEN env"`
	Group           string `short:"g" long:"group" description:"Gitlab project group, defaults to GITLAB_GROUP env"`
	Project         string `short:"p" long:"project" description:"Gitlab project, defaults to GITLAB_PROJECT env"`
	Branch          string `short:"b" long:"branch" description:"Gitlab branch, defaults to GITLAB_BRANCH env, or 'main'"`
	Job             string `short:"j" long:"job" description:"Gitlab job, defaults to GITLAB_JOB env"`
	Force           bool   `short:"f" long:"force" description:"Force download, even if not changed"`
	Verbose         bool   `short:"v" long:"verbose" description:"Be verbose"`
	Debug           bool   `short:"d" long:"debug" description:"Debug messages"`
	ShowVersion     bool   `long:"version" description:"Show version information and exit"`
	RequestTimeout  int    `long:"request-timeout" description:"Request timeout in seconds, for JSON API http requests"`
	DownloadTimeout int    `long:"download-timeout" description:"Download timeout in seconds, for artifact download http requests"`
}
