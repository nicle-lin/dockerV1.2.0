package execdrivers

import (
	"fmt"
	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver"
	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver/lxc"
	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver/native"
	"github.com/nicle-lin/dockerV1.2.0/docker/pkg/sysinfo"
	"path"
)

func NewDriver(name, root, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
	switch name {
	case "lxc":
		// we want to give the lxc driver the full docker root because it needs
		// to access and write config and template files in /var/lib/docker/containers/*
		// to be backwards compatible
		return lxc.NewDriver(root, initPath, sysInfo.AppArmor)
	case "native":
		return native.NewDriver(path.Join(root, "execdriver", "native"), initPath)
	}
	return nil, fmt.Errorf("unknown exec driver %s", name)
}
