package webhooks

type Events struct {
	domainNameStr  string
	eventSlugsStrs []string
}

func NewEvents(domainName string, eventSlugs []string) Events {
	return Events{
		domainNameStr:  domainName,
		eventSlugsStrs: eventSlugs}
}

func (events *Events) DomainName() string {
	return events.domainNameStr
}

func (events *Events) EventSlugs() []string {
	return events.eventSlugsStrs
}
