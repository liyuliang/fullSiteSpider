package config

type web struct {
	PORT    string
}

func Web() web {

	return getConfig().WEB
}
