// +build !linux

package nodes

import (
	"errors"

	"github.com/nicle-lin/dockerV1.2.0/libcontainer/devices"
)

func CreateDeviceNodes(rootfs string, nodesToCreate []*devices.Device) error {
	return errors.New("Unsupported method")
}
