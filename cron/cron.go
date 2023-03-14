package cron

import (
	"github.com/yangkaiyue/gin-exp/global"
	"time"
)

func CrontabJob() {

	go func() {
		t := time.NewTicker(time.Second * 30)
		for range t.C {
			SayAlive()
		}
	}()
}

func SayAlive() {
	global.Logger.Info("I'm Alive")
}
