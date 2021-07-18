package mullvad

import (
	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/utils"
)

func (m *Mullvad) GetOpenVPNConnection(selection configuration.ServerSelection) (
	connection models.OpenVPNConnection, err error) {
	var port uint16 = 1194
	protocol := constants.UDP
	if selection.TCP {
		port = 443
		protocol = constants.TCP
	}

	if selection.CustomPort > 0 {
		port = selection.CustomPort
	}

	servers, err := m.filterServers(selection)
	if err != nil {
		return connection, err
	}

	var connections []models.OpenVPNConnection
	for _, server := range servers {
		for _, IP := range server.IPs {
			connection := models.OpenVPNConnection{
				IP:       IP,
				Port:     port,
				Protocol: protocol,
			}
			connections = append(connections, connection)
		}
	}

	if selection.TargetIP != nil {
		return utils.GetTargetIPConnection(connections, selection.TargetIP)
	}

	return utils.PickRandomConnection(connections, m.randSource), nil
}
