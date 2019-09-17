package service

import (
	"github.com/liyuliang/queue-services"
)

func Start() {
	services.Service().Start(true)
}
