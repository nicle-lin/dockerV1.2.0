package libcontainer

import (
	"github.com/nicle-lin/dockerV1.2.0/libcontainer/cgroups"
	"github.com/nicle-lin/dockerV1.2.0/libcontainer/network"
)

type ContainerStats struct {
	NetworkStats *network.NetworkStats `json:"network_stats,omitempty"`
	CgroupStats  *cgroups.Stats        `json:"cgroup_stats,omitempty"`
}
