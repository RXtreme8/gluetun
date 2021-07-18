package protonvpn

import (
	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/utils"
)

func (p *Protonvpn) GetOpenVPNConnection(selection configuration.ServerSelection) (
	connection models.OpenVPNConnection, err error) {
	protocol := constants.UDP
	if selection.TCP {
		protocol = constants.TCP
	}

	port, err := getPort(selection.TCP, selection.CustomPort)
	if err != nil {
		return connection, err
	}

	servers, err := p.filterServers(selection)
	if err != nil {
		return connection, err
	}

	connections := make([]models.OpenVPNConnection, len(servers))
	for i := range servers {
		connections[i] = models.OpenVPNConnection{
			IP:       servers[i].EntryIP,
			Port:     port,
			Protocol: protocol,
		}
	}

	if selection.TargetIP != nil {
		return utils.GetTargetIPConnection(connections, selection.TargetIP)
	}

	return utils.PickRandomConnection(connections, p.randSource), nil
}
