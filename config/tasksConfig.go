package config

import "net/url"

type Task struct {
	Url           string
	Keywords      []string
	EmailTo       string
	TitleSelector string
	HrefsSelector string
}

func Tasks() map[string]Task {
	return getConfig().Tasks
}

func (t Task) Domain() (domain string) {
	u, err := url.Parse(t.Url)
	if err == nil {
		domain = u.Hostname()
	}
	return domain
}
