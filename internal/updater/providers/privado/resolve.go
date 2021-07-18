package privado

import (
	"context"
	"net"
	"time"

	"github.com/rxtreme8/gluetun/internal/updater/resolver"
)

func resolveHosts(ctx context.Context, presolver resolver.Parallel,
	hosts []string, minServers int) (hostToIPs map[string][]net.IP,
	warnings []string, err error) {
	const (
		maxFailRatio = 0.1
		maxDuration  = 3 * time.Second
		maxNoNew     = 1
		maxFails     = 2
	)
	settings := resolver.ParallelSettings{
		MaxFailRatio: maxFailRatio,
		MinFound:     minServers,
		Repeat: resolver.RepeatSettings{
			MaxDuration: maxDuration,
			MaxNoNew:    maxNoNew,
			MaxFails:    maxFails,
			SortIPs:     true,
		},
	}
	return presolver.Resolve(ctx, hosts, settings)
}
