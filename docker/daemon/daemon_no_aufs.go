// +build exclude_graphdriver_aufs

package daemon

import (
	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/graphdriver"
)

func migrateIfAufs(driver graphdriver.Driver, root string) error {
	return nil
}
