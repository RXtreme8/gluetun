package protonvpn

import (
	"strings"

	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/utils"
)

func (p *Protonvpn) filterServers(selection configuration.ServerSelection) (
	servers []models.ProtonvpnServer, err error) {
	for _, server := range p.servers {
		switch {
		case
			utils.FilterByPossibilities(server.Country, selection.Countries),
			utils.FilterByPossibilities(server.Region, selection.Regions),
			utils.FilterByPossibilities(server.City, selection.Cities),
			utils.FilterByPossibilities(server.Hostname, selection.Hostnames),
			utils.FilterByPossibilities(server.Name, selection.Names),
			selection.FreeOnly && !strings.Contains(strings.ToLower(server.Name), "free"):
		default:
			servers = append(servers, server)
		}
	}

	if len(servers) == 0 {
		return nil, utils.NoServerFoundError(selection)
	}

	return servers, nil
}
