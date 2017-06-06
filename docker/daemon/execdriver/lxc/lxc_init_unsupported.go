// +build !linux

package lxc

import "github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver"

func setHostname(hostname string) error {
	panic("Not supported on darwin")
}

func finalizeNamespace(args *execdriver.InitArgs) error {
	panic("Not supported on darwin")
}
