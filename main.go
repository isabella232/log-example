package main

import (
	"os"
	"time"

	"github.com/leodotcloud/log"
	"github.com/leodotcloud/log-example/one"
	"github.com/leodotcloud/log-example/two"
	"github.com/leodotcloud/log/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Turn on debug logging",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	server.StartServerWithDefaults()

	if c.Bool("debug") {
		log.SetLevelString("debug")
	}

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
