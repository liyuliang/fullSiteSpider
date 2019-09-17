package cli

import (
	"flag"
)

var commands = []command{}

func init() {

	commands = append(commands, command{
		Name: "config",
		value:flag.String("c", "./config/config.toml", "toml type config file path"),
	})

	flag.Parse()
}

func Parse(cliKey string) (val string) {

	for _, _command := range commands {
		if _command.Name == cliKey {
			val = *_command.value
		}
	}
	return val
}
