package ipvanish

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Ipvanish struct {
	servers    []models.IpvanishServer
	randSource rand.Source
}

func New(servers []models.IpvanishServer, randSource rand.Source) *Ipvanish {
	return &Ipvanish{
		servers:    servers,
		randSource: randSource,
	}
}
