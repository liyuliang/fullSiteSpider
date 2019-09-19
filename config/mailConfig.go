package config

type mail struct {
	Account  string
	Password string
	Port     string
	SmtpHost string
}

func Mail() mail {
	return getConfig().Mail
}
