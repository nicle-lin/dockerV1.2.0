package main

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/nicle-lin/dockerV1.2.0/libcontainer"
	_ "github.com/nicle-lin/dockerV1.2.0/libcontainer/nsenter"
)

var initCommand = cli.Command{
	Name:  "init",
	Usage: "runs the init process inside the namespace",
	Action: func(context *cli.Context) {
		logrus.SetLevel(logrus.DebugLevel)
		runtime.GOMAXPROCS(1)
		runtime.LockOSThread()
		factory, err := libcontainer.New("")
		if err != nil {
			fatal(err)
		}
		if err := factory.StartInitialization(); err != nil {
			fatal(err)
		}
		panic("This line should never been executed")
	},
}
