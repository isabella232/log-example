# Example using github.com/leodotcloud/log package


```go
package main

import (
	"time"
	log "github.com/leodotcloud/log"
)

func main() {
	num := 0
	for {
		time.Sleep(time.Millisecond * 1000)
		log.Infof("%v :: Number: %v", num)
		num++
		time.Sleep(time.Millisecond * 1000)
		log.Errorf("%v :: Number: %v", num)
		num++
		time.Sleep(time.Millisecond * 1000)
		log.Debugf("%v :: Number: %v", num)
		num++
	}
}
```

See `main.go` for a more complex example.

Example using a container:

```bash
docker run --name=example -itd leodotcloud/log-example:dev
docker logs -f example
```

To change the log level dynamically:

``` bash
docker exec -it example setLogLevelDebug
docker exec -it example getLogLevel
docker exec -it example setLogLevelInfo
```
