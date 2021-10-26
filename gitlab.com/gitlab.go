package gitlab

import "time"

// EventPush is defined at https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html#push-events
type EventPush struct {
	ObjectKind        string     `json:"object_kind,omitempty"`
	EventName         string     `json:"event_name,omitempty"`
	Before            string     `json:"before,omitempty"`
	After             string     `json:"after,omitempty"`
	Ref               string     `json:"ref,omitempty"`
	CheckoutSha       string     `json:"checkout_sha,omitempty"`
	UserId            string     `json:"user_id"`
	UserName          string     `json:"user_name,omitempty"`
	UserUsername      string     `json:"user_username,omitempty"`
	UserEmail         string     `json:"user_email,omitempty"`
	UserAvatar        string     `json:"user_avatar,omitempty"`
	ProjectId         string     `json:"project_id"`
	Project           Project    `json:"project,omitempty"`
	Repository        Repository `json:"repository,omitempty"`
	Commits           []Commit   `json:"commits,omitempty"`
	TotalCommitsCount int        `json:"total_commits_count"` // include `0`
}

type EventMergeRequest struct {
	ObjectKind string     `json:"object_kind,omitempty"`
	EventName  string     `json:"event_name,omitempty"`
	User       User       `json:"user,omitempty"`
	Project    Project    `json:"project,omitempty"`
	Repository Repository `json:"repository,omitempty"`
}

// User is used by `merge_request` event
type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Email     string `json:"email,omitempty"`
}

// Repository is used by `push` event
type Repository struct {
	Name            string `json:"name,omitempty"`
	URL             string `json:"url,omitempty"`
	Description     string `json:"description,omitempty"`
	Homepage        string `json:"homepage,omitempty"`
	GitHttpUrl      string `json:"git_http_url,omitempty"`
	GitSshUrl       string `json:"git_ssh_url,omitempty"`
	VisibilityLevel int    `json:"visibility_level"` // include `0`
}

// Project is used by `push` event
type Project struct {
	Id                int    `json:"id"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	WebUrl            string `json:"web_url,omitempty"`
	AvatarUrl         string `json:"avatar_url,omitempty"`
	GitSshUrl         string `json:"git_ssh_url,omitempty"`
	GitHttpUrl        string `json:"git_http_url,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	VisibilityLevel   int    `json:"visibility_level"` // include `0`
	PathWithNamespace string `json:"path_with_namespace,omitempty"`
	DefaultBranch     string `json:"default_branch,omitempty"`
	Homepage          string `json:"homepage,omitempty"`
	Url               string `json:"url,omitempty"`
	SshUrl            string `json:"ssh_url,omitempty"`
	HttpUrl           string `json:"http_url,omitempty"`
}

// Commit is used by `push` event
type Commit struct {
	Id        string    `json:"id"`
	Message   string    `json:"message,omitempty"`
	Title     string    `json:"title,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Url       string    `json:"url,omitempty"`
	Author    Author    `json:"author,omitempty"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
}

// Author is used by `push` event
type Author struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
