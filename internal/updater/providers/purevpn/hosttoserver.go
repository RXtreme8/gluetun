package purevpn

import (
	"net"

	"github.com/rxtreme8/gluetun/internal/models"
)

type hostToServer map[string]models.PurevpnServer

func (hts hostToServer) add(host string, tcp, udp bool) {
	server, ok := hts[host]
	if !ok {
		server.Hostname = host
	}
	if tcp {
		server.TCP = true
	}
	if udp {
		server.UDP = true
	}
	hts[host] = server
}

func (hts hostToServer) toHostsSlice() (hosts []string) {
	hosts = make([]string, 0, len(hts))
	for host := range hts {
		hosts = append(hosts, host)
	}
	return hosts
}

func (hts hostToServer) adaptWithIPs(hostToIPs map[string][]net.IP) {
	for host, IPs := range hostToIPs {
		server := hts[host]
		server.IPs = IPs
		hts[host] = server
	}
	for host, server := range hts {
		if len(server.IPs) == 0 {
			delete(hts, host)
		}
	}
}

func (hts hostToServer) toServersSlice() (servers []models.PurevpnServer) {
	servers = make([]models.PurevpnServer, 0, len(hts))
	for _, server := range hts {
		servers = append(servers, server)
	}
	return servers
}
