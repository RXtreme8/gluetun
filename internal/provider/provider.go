// Package provider defines interfaces to interact with each VPN provider.
package provider

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/firewall"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/cyberghost"
	"github.com/rxtreme8/gluetun/internal/provider/fastestvpn"
	"github.com/rxtreme8/gluetun/internal/provider/hidemyass"
	"github.com/rxtreme8/gluetun/internal/provider/ipvanish"
	"github.com/rxtreme8/gluetun/internal/provider/ivpn"
	"github.com/rxtreme8/gluetun/internal/provider/mullvad"
	"github.com/rxtreme8/gluetun/internal/provider/nordvpn"
	"github.com/rxtreme8/gluetun/internal/provider/privado"
	"github.com/rxtreme8/gluetun/internal/provider/privateinternetaccess"
	"github.com/rxtreme8/gluetun/internal/provider/privatevpn"
	"github.com/rxtreme8/gluetun/internal/provider/protonvpn"
	"github.com/rxtreme8/gluetun/internal/provider/purevpn"
	"github.com/rxtreme8/gluetun/internal/provider/surfshark"
	"github.com/rxtreme8/gluetun/internal/provider/torguard"
	"github.com/rxtreme8/gluetun/internal/provider/vpnunlimited"
	"github.com/rxtreme8/gluetun/internal/provider/vyprvpn"
	"github.com/rxtreme8/gluetun/internal/provider/windscribe"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/os"
)

// Provider contains methods to read and modify the openvpn configuration to connect as a client.
type Provider interface {
	GetOpenVPNConnection(selection configuration.ServerSelection) (connection models.OpenVPNConnection, err error)
	BuildConf(connection models.OpenVPNConnection, username string, settings configuration.OpenVPN) (lines []string)
	PortForward(ctx context.Context, client *http.Client,
		openFile os.OpenFileFunc, pfLogger logging.Logger, gateway net.IP, fw firewall.Configurator,
		syncState func(port uint16) (pfFilepath string))
}

func New(provider string, allServers models.AllServers, timeNow func() time.Time) Provider {
	randSource := rand.NewSource(timeNow().UnixNano())
	switch provider {
	case constants.Cyberghost:
		return cyberghost.New(allServers.Cyberghost.Servers, randSource)
	case constants.Fastestvpn:
		return fastestvpn.New(allServers.Fastestvpn.Servers, randSource)
	case constants.HideMyAss:
		return hidemyass.New(allServers.HideMyAss.Servers, randSource)
	case constants.Ipvanish:
		return ipvanish.New(allServers.Ipvanish.Servers, randSource)
	case constants.Ivpn:
		return ivpn.New(allServers.Ivpn.Servers, randSource)
	case constants.Mullvad:
		return mullvad.New(allServers.Mullvad.Servers, randSource)
	case constants.Nordvpn:
		return nordvpn.New(allServers.Nordvpn.Servers, randSource)
	case constants.Privado:
		return privado.New(allServers.Privado.Servers, randSource)
	case constants.PrivateInternetAccess:
		return privateinternetaccess.New(allServers.Pia.Servers, randSource, timeNow)
	case constants.Privatevpn:
		return privatevpn.New(allServers.Privatevpn.Servers, randSource)
	case constants.Protonvpn:
		return protonvpn.New(allServers.Protonvpn.Servers, randSource)
	case constants.Purevpn:
		return purevpn.New(allServers.Purevpn.Servers, randSource)
	case constants.Surfshark:
		return surfshark.New(allServers.Surfshark.Servers, randSource)
	case constants.Torguard:
		return torguard.New(allServers.Torguard.Servers, randSource)
	case constants.VPNUnlimited:
		return vpnunlimited.New(allServers.VPNUnlimited.Servers, randSource)
	case constants.Vyprvpn:
		return vyprvpn.New(allServers.Vyprvpn.Servers, randSource)
	case constants.Windscribe:
		return windscribe.New(allServers.Windscribe.Servers, randSource)
	default:
		return nil // should never occur
	}
}
