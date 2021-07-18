package privado

import (
	"math/rand"

	"github.com/rxtreme8/gluetun/internal/models"
)

type Privado struct {
	servers    []models.PrivadoServer
	randSource rand.Source
}

func New(servers []models.PrivadoServer, randSource rand.Source) *Privado {
	return &Privado{
		servers:    servers,
		randSource: randSource,
	}
}
