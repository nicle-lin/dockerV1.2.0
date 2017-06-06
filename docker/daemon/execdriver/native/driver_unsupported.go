// +build !linux

package native

import (
	"fmt"

	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/execdriver"
)

func NewDriver(root, initPath string) (execdriver.Driver, error) {
	return nil, fmt.Errorf("native driver not supported on non-linux")
}
