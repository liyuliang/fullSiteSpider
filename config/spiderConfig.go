package config

type spider struct {
	KEYWORDS []string
}

func Spider() spider {

	return getConfig().SPIDER
}
