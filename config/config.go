package config

import (
	"github.com/BurntSushi/toml"
	"github.com/liyuliang/utils/path"
	"os"
)

const appConfigFile string = "config.toml"

type appConfig struct {
	VERSION string
	WEB     web
}

var _appConfigFilePath string
var _appConfig appConfig

func initConfig() {

	_, err := toml.DecodeFile(appConfigFilePath(), &_appConfig)

	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
}

func appConfigFilePath() string {

	if _appConfigFilePath == "" {
		_appConfigFilePath = path.CURRENT_DIR() + "/../" + appConfigFile
	}
	return _appConfigFilePath
}

func SetAppConfigFilePath(path string) {
	println("reading the config file:", path)

	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		println("file exist")
		_appConfigFilePath = path
	} else {
		println("file not exist!")
	}
}

func getConfig() appConfig {
	if _appConfig.VERSION == "" {
		initConfig()
	}
	return _appConfig
}

func reConfig() appConfig {

	initConfig()
	return _appConfig
}
