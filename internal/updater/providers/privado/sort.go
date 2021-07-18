package privado

import (
	"sort"

	"github.com/rxtreme8/gluetun/internal/models"
)

func sortServers(servers []models.PrivadoServer) {
	sort.Slice(servers, func(i, j int) bool {
		if servers[i].Country == servers[j].Country {
			if servers[i].Region == servers[j].Region {
				if servers[i].City == servers[j].City {
					return servers[i].Hostname < servers[j].Hostname
				}
				return servers[i].City < servers[j].City
			}
			return servers[i].Region < servers[j].Region
		}
		return servers[i].Country < servers[j].Country
	})
}
