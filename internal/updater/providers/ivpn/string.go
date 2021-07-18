package ivpn

import "github.com/rxtreme8/gluetun/internal/models"

func Stringify(servers []models.IvpnServer) (s string) {
	s = "func IvpnServers() []models.IvpnServer {\n"
	s += "	return []models.IvpnServer{\n"
	for _, server := range servers {
		s += "		" + server.String() + ",\n"
	}
	s += "	}\n"
	s += "}"
	return s
}
