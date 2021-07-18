package nordvpn

import "github.com/rxtreme8/gluetun/internal/models"

func Stringify(servers []models.NordvpnServer) (s string) {
	s = "func NordvpnServers() []models.NordvpnServer {\n"
	s += "	return []models.NordvpnServer{\n"
	for _, server := range servers {
		s += "		" + server.String() + ",\n"
	}
	s += "	}\n"
	s += "}"
	return s
}
