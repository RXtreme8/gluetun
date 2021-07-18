package privatevpn

import "github.com/rxtreme8/gluetun/internal/models"

func Stringify(servers []models.PrivatevpnServer) (s string) {
	s = "func PrivatevpnServers() []models.PrivatevpnServer {\n"
	s += "	return []models.PrivatevpnServer{\n"
	for _, server := range servers {
		s += "		" + server.String() + ",\n"
	}
	s += "	}\n"
	s += "}"
	return s
}
