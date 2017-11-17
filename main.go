package main

import (
	"time"

	"github.com/leodotcloud/log"
	"github.com/leodotcloud/log-example/one"
	"github.com/leodotcloud/log-example/two"
	"github.com/leodotcloud/log/server"
)

func main() {

	server.StartServerWithDefaults()

	one.Watch()
	two.Watch()

	id := "main"
	num := 0
	for {
		time.Sleep(time.Millisecond * 1000)
		log.Infof("%v :: Number: %v", id, num)
		num++
		time.Sleep(time.Millisecond * 1000)
		log.Errorf("%v :: Number: %v", id, num)
		num++
		time.Sleep(time.Millisecond * 1000)
		log.Debugf("%v :: Number: %v", id, num)
		num++
	}
}
