package config

import (
	"github.com/liyuliang/utils/cli"
	"log"
)

func Init() {

	param := cli.GetParam(1)

	if param.IsExist() {

		configFilePath := param.ToString()
		SetAppConfigFilePath(configFilePath)
	}
}
