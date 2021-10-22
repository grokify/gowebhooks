package webhooks

type Events struct {
	serviceSlugStr string
	domainNameStr  string
	eventSlugsStrs []string
}

func NewEvents(serviceSlug, domainName string, eventSlugs []string) Events {
	return Events{
		serviceSlugStr: serviceSlug,
		domainNameStr:  domainName,
		eventSlugsStrs: eventSlugs}
}

func (events *Events) DomainName() string {
	return events.domainNameStr
}

func (events *Events) EventSlugs() []string {
	return events.eventSlugsStrs
}

func (events *Events) ServiceSlug() string {
	return events.serviceSlugStr
}
