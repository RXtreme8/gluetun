package vyprvpn

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Vyprvpn struct {
	servers    []models.VyprvpnServer
	randSource rand.Source
}

func New(servers []models.VyprvpnServer, randSource rand.Source) *Vyprvpn {
	return &Vyprvpn{
		servers:    servers,
		randSource: randSource,
	}
}
