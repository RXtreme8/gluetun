package privateinternetaccess

import (
	"math/rand"
	"time"

	"github.com/rxtreme8/gluetun/internal/models"
)

type PIA struct {
	servers      []models.PIAServer
	randSource   rand.Source
	timeNow      func() time.Time
	activeServer models.PIAServer
}

func New(servers []models.PIAServer, randSource rand.Source, timeNow func() time.Time) *PIA {
	return &PIA{
		servers:    servers,
		timeNow:    timeNow,
		randSource: randSource,
	}
}
