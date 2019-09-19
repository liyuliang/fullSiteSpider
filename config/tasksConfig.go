package config

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
