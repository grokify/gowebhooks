package gitlab

import (
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
