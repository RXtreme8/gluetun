package ivpn

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Ivpn struct {
	servers    []models.IvpnServer
	randSource rand.Source
}

func New(servers []models.IvpnServer, randSource rand.Source) *Ivpn {
	return &Ivpn{
		servers:    servers,
		randSource: randSource,
	}
}
