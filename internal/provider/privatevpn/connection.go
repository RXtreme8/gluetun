package privatevpn

import (
	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/utils"
)

func (p *Privatevpn) GetOpenVPNConnection(selection configuration.ServerSelection) (
	connection models.OpenVPNConnection, err error) {
	protocol := constants.UDP
	var port uint16 = 1194
	if selection.TCP {
		protocol = constants.TCP
		port = 443
	}

	servers, err := p.filterServers(selection)
	if err != nil {
		return connection, err
	}

	var connections []models.OpenVPNConnection
	for _, server := range servers {
		for _, ip := range server.IPs {
			connection := models.OpenVPNConnection{
				IP:       ip,
				Port:     port,
				Protocol: protocol,
			}
			connections = append(connections, connection)
		}
	}

	if selection.TargetIP != nil {
		return utils.GetTargetIPConnection(connections, selection.TargetIP)
	}

	return utils.PickRandomConnection(connections, p.randSource), nil
}
