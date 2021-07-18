package fastestvpn

import "github.com/rxtreme8/gluetun/internal/models"

func Stringify(servers []models.FastestvpnServer) (s string) {
	s = "func FastestvpnServers() []models.FastestvpnServer {\n"
	s += "	return []models.FastestvpnServer{\n"
	for _, server := range servers {
		s += "		" + server.String() + ",\n"
	}
	s += "	}\n"
	s += "}"
	return s
}
