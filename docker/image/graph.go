package image

import (
	"github.com/nicle-lin/dockerV1.2.0/docker/daemon/graphdriver"
)

type Graph interface {
	Get(id string) (*Image, error)
	ImageRoot(id string) string
	Driver() graphdriver.Driver
}
