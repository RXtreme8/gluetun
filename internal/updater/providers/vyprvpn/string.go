package vyprvpn

import "github.com/rxtreme8/gluetun/internal/models"

func Stringify(servers []models.VyprvpnServer) (s string) {
	s = "func VyprvpnServers() []models.VyprvpnServer {\n"
	s += "	return []models.VyprvpnServer{\n"
	for _, server := range servers {
		s += "		" + server.String() + ",\n"
	}
	s += "	}\n"
	s += "}"
	return s
}
