package libcontainer

import "github.com/nicle-lin/dockerV1.2.0/libcontainer/cgroups"

type Stats struct {
	Interfaces  []*NetworkInterface
	CgroupStats *cgroups.Stats
}
