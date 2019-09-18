package config

import "testing"

func TestRedis(t *testing.T) {

	if 6379 != Redis().Ports[0] {
		t.Error("redis main port is not 6379")
	}
}
