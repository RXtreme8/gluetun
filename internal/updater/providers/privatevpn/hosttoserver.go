package privatevpn

import (
	"net"

	"github.com/rxtreme8/gluetun/internal/models"
)

type hostToServer map[string]models.PrivatevpnServer

// TODO check if server supports TCP and UDP.
func (hts hostToServer) add(host, country, city string) {
	server, ok := hts[host]
	if ok {
		return
	}
	server.Hostname = host
	server.Country = country
	server.City = city
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

func (hts hostToServer) toServersSlice() (servers []models.PrivatevpnServer) {
	servers = make([]models.PrivatevpnServer, 0, len(hts))
	for _, server := range hts {
		servers = append(servers, server)
	}
	return servers
}
