package devices

import (
	"github.com/nicle-lin/dockerV1.2.0/libcontainer/configs"
)

// TODO Windows. This can be factored out further - Devices are not supported
// by Windows Containers.

func DeviceFromPath(path, permissions string) (*configs.Device, error) {
	return nil, nil
}

func HostDevices() ([]*configs.Device, error) {
	return nil, nil
}
