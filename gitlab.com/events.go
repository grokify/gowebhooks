package gitlab

import (
	"strings"

	webhooks "github.com/grokify/webhook-directory"
)

const (
	DomainName = "gitlab.com"
	EventSlugs = "comment-on-commit,comment-onissue,comment-on-merge-request,issue,merge-request,push,release,tag-push"
)

func NewEvents() webhooks.Events {
	return webhooks.NewEvents(
		DomainName,
		strings.Split(EventSlugs, ","))
}
