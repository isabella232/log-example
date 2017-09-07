package main

import (
	"time"

	log "github.com/leodotcloud/log"
	"github.com/leodotcloud/log-example/one"
	"github.com/leodotcloud/log-example/server"
	"github.com/leodotcloud/log-example/two"
)

func main() {

	go func() {
		s := server.Server{}
		s.ListenAndServe("localhost:12345")
	}()

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
