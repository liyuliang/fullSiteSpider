package config

type task struct {
	Url      string
	Keywords []string
	EmailTo  string
}

func Tasks() map[string]task {
	return getConfig().Tasks
}
