package vyprvpn

import (
	"context"
	"net"
	"net/http"

	"github.com/rxtreme8/gluetun/internal/firewall"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/os"
)

func (v *Vyprvpn) PortForward(ctx context.Context, clienv *http.Client,
	openFile os.OpenFileFunc, pfLogger logging.Logger, gateway net.IP,
	fw firewall.Configurator, syncState func(port uint16) (pfFilepath string)) {
	panic("port forwarding is not supported for Vyprvpn")
}
