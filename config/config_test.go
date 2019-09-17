package config

import (
	"testing"
)

func Test_reConfig(t *testing.T) {

	getConfig()

	if _appConfig.VERSION == "" {
		t.Error("read config file failed")
	}

	if _appConfig.WEB.PORT != "8081" {
		t.Error("read config http web port is not 8081: ", _appConfig.WEB.PORT)
	}
}
