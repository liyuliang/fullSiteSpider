package config

import (
	"github.com/liyuliang/utils/format"
	"time"
)

type redis struct {
	Name                string
	Server              string
	Ports               []int
	ConnectionMaxSecond int`toml:"connection_max_second"`
	MaxIdle             int`toml:"max_idle"`
	DBNum               int`toml:"db_number"`
	Password            string
	Enabled             bool
}

func Redis() redis {
	return getConfig().Redis
}

func (_redis redis) MainPort() string {
	return format.IntToStr(_redis.Ports[0])
}

func (_redis redis) Link() string {
	return _redis.Server + ":" + _redis.MainPort()
}

func (_redis redis) ConnectionMax() time.Duration {
	return format.IntToTimeSecond(_redis.ConnectionMaxSecond)
}

func (_redis redis) DBNumber() int {
	return _redis.DBNum
}
