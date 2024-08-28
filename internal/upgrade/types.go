package upgrade

import "time"

type JobListCommit struct {
	AuthorEmail string    `json:"author_email"`
	AuthorName  string    `json:"author_name"`
	CreatedAt   time.Time `json:"created_at"`
	Id          string    `json:"id"`
	Message     string    `json:"message"`
	ShortId     string    `json:"short_id"`
	Title       string    `json:"title"`
}

type JobListArtifact struct {
	FileType   string `json:"file_type"`
	Size       int    `json:"size"`
	Filename   string `json:"filename"`
	FileFormat string `json:"file_format"`
}

type JobListPipeline struct {
	Id        int    `json:"id"`
	ProjectId int    `json:"project_id"`
	Ref       string `json:"ref"`
	Sha       string `json:"sha"`
	Status    string `json:"status"`
}

type JobListRunner struct {
	Id          int         `json:"id"`
	Description string      `json:"description"`
	IpAddress   interface{} `json:"ip_address"`
	Active      bool        `json:"active"`
	Paused      bool        `json:"paused"`
	IsShared    bool        `json:"is_shared"`
	RunnerType  string      `json:"runner_type"`
	Name        interface{} `json:"name"`
	Online      bool        `json:"online"`
	Status      string      `json:"status"`
}

type JobListRunnerManager struct {
	Id           int       `json:"id"`
	SystemId     string    `json:"system_id"`
	Version      string    `json:"version"`
	Revision     string    `json:"revision"`
	Platform     string    `json:"platform"`
	Architecture string    `json:"architecture"`
	CreatedAt    time.Time `json:"created_at"`
	ContactedAt  time.Time `json:"contacted_at"`
	IpAddress    string    `json:"ip_address"`
	Status       string    `json:"status"`
}

type JobListUser struct {
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Username     string      `json:"username"`
	State        string      `json:"state"`
	AvatarUrl    string      `json:"avatar_url"`
	WebUrl       string      `json:"web_url"`
	CreatedAt    time.Time   `json:"created_at"`
	Bio          interface{} `json:"bio"`
	Location     interface{} `json:"location"`
	PublicEmail  string      `json:"public_email"`
	Skype        string      `json:"skype"`
	Linkedin     string      `json:"linkedin"`
	Twitter      string      `json:"twitter"`
	WebsiteUrl   string      `json:"website_url"`
	Organization string      `json:"organization"`
}

type JobListArtifactsFile struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type JobListItem struct {
	Commit JobListCommit `json:"commit"`
	// Coverage       ???  `json:"coverage"`
	Archived          bool                 `json:"archived"`
	AllowFailure      bool                 `json:"allow_failure"`
	CreatedAt         time.Time            `json:"created_at"`
	StartedAt         time.Time            `json:"started_at"`
	FinishedAt        time.Time            `json:"finished_at"`
	ErasedAt          interface{}          `json:"erased_at"`
	Duration          float64              `json:"duration"`
	QueuedDuration    float64              `json:"queued_duration"`
	ArtifactsFile     JobListArtifactsFile `json:"artifacts_file"`
	Artifacts         []JobListArtifact    `json:"artifacts"`
	ArtifactsExpireAt time.Time            `json:"artifacts_expire_at"`
	TagList           []string             `json:"tag_list"`
	Id                int                  `json:"id"`
	Name              string               `json:"name"`
	Pipeline          JobListPipeline      `json:"pipeline"`
	Ref               string               `json:"ref"`
	Runner            JobListRunner        `json:"runner"`
	RunnerManager     JobListRunnerManager `json:"runner_manager"`
	Stage             string               `json:"stage"`
	Status            string               `json:"status"`
	FailureReason     string               `json:"failure_reason"`
	Tag               bool                 `json:"tag"`
	WebUrl            string               `json:"web_url"`
	Project           struct {
		CiJobTokenScopeEnabled bool `json:"ci_job_token_scope_enabled"`
	} `json:"project"`
	User JobListUser `json:"user"`
}
