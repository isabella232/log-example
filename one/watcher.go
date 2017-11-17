package one

import (
	"time"

	"github.com/leodotcloud/log"
)

// Watch ...
func Watch() {

	go func(id string) {
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
	}("one")
}
