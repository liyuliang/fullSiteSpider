package config

type task struct {
	Homes []string
}

func Task() task {
	return getConfig().TASK
}
