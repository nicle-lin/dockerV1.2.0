package main

import (
	_ "github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver/lxc"
	_ "github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver/native"
	"github.com/nicle-lin/dockerV1.2.0/docker/reexec"
)

func main() {
	// Running in init mode
	reexec.Init()
}
