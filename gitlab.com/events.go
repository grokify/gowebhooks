package gitlab

import (
	"embed"
	"strings"

	"github.com/grokify/gowebhooks"
)

const (
	ServiceSlug = "gitlab"
	DomainName  = "gitlab.com"
	EventSlugs  = "comment-on-commit,comment-onissue,comment-on-merge-request,issue,merge-request,push,release,tag-push"
)

func NewEvents() gowebhooks.Events {
	return gowebhooks.NewEvents(
		ServiceSlug,
		DomainName,
		strings.Split(EventSlugs, ","))
}

//go:embed event-example_push_slack.json
//go:embed event-example_push_slack.txt
var f embed.FS

// ReadExampleFile reads an embedded example file.
func ReadExampleFile(filename string) ([]byte, error) {
	return f.ReadFile(filename)
}
