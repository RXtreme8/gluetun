package windscribe

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Windscribe struct {
	servers    []models.WindscribeServer
	randSource rand.Source
}

func New(servers []models.WindscribeServer, randSource rand.Source) *Windscribe {
	return &Windscribe{
		servers:    servers,
		randSource: randSource,
	}
}
