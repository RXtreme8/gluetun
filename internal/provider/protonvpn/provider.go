package protonvpn

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Protonvpn struct {
	servers    []models.ProtonvpnServer
	randSource rand.Source
}

func New(servers []models.ProtonvpnServer, randSource rand.Source) *Protonvpn {
	return &Protonvpn{
		servers:    servers,
		randSource: randSource,
	}
}
