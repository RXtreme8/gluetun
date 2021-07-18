package pia

import (
	"net"

	"github.com/rxtreme8/gluetun/internal/models"
)

type nameToServer map[string]models.PIAServer

func (nts nameToServer) add(name, hostname, region string,
	tcp, udp, portForward bool, ip net.IP) (change bool) {
	server, ok := nts[name]
	if !ok {
		change = true
		server.ServerName = name
		server.Hostname = hostname
		server.Region = region
		server.PortForward = portForward
	}

	if !server.TCP && tcp {
		change = true
		server.TCP = tcp
	}
	if !server.UDP && udp {
		change = true
		server.UDP = udp
	}

	ipFound := false
	for _, existingIP := range server.IPs {
		if ip.Equal(existingIP) {
			ipFound = true
			break
		}
	}

	if !ipFound {
		change = true
		server.IPs = append(server.IPs, ip)
	}

	nts[name] = server

	return change
}

func (nts nameToServer) toServersSlice() (servers []models.PIAServer) {
	servers = make([]models.PIAServer, 0, len(nts))
	for _, server := range nts {
		servers = append(servers, server)
	}
	return servers
}
