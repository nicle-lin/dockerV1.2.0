// +build linux

package daemon

import "github.com/nicle-lin/dockerV1.2.0/libcontainer/selinux"

func selinuxSetDisabled() {
	selinux.SetDisabled()
}

func selinuxFreeLxcContexts(label string) {
	selinux.FreeLxcContexts(label)
}
