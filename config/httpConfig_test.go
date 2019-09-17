package config

import (
	"testing"
)

func Test_getInfo(t *testing.T) {

	if Web().PORT == "" {
		t.Error("can not get http port from toml config file")
	}
}
