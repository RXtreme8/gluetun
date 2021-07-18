package nordvpn

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Nordvpn struct {
	servers    []models.NordvpnServer
	randSource rand.Source
}

func New(servers []models.NordvpnServer, randSource rand.Source) *Nordvpn {
	return &Nordvpn{
		servers:    servers,
		randSource: randSource,
	}
}
