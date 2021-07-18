package surfshark

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Surfshark struct {
	servers    []models.SurfsharkServer
	randSource rand.Source
}

func New(servers []models.SurfsharkServer, randSource rand.Source) *Surfshark {
	return &Surfshark{
		servers:    servers,
		randSource: randSource,
	}
}
