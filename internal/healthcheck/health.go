// Package healthcheck defines the client and server side of the built in healthcheck.
package healthcheck

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"
)

func (s *server) runHealthcheckLoop(ctx context.Context, done chan<- struct{}) {
	defer close(done)

	s.openvpn.healthyTimer = time.NewTimer(defaultOpenvpnHealthyWaitTime)

	for {
		previousErr := s.handler.getErr()

		err := healthCheck(ctx, s.resolver)
		s.handler.setErr(err)

		if previousErr != nil && err == nil {
			s.logger.Info("healthy!")
			s.openvpn.healthyTimer.Stop()
			s.openvpn.healthyWaitTime = defaultOpenvpnHealthyWaitTime
		} else if previousErr == nil && err != nil {
			s.logger.Info("unhealthy: " + err.Error())
			s.openvpn.healthyTimer = time.NewTimer(s.openvpn.healthyWaitTime)
		}

		if err != nil { // try again after 1 second
			timer := time.NewTimer(time.Second)
			select {
			case <-ctx.Done():
				if !timer.Stop() {
					<-timer.C
				}
				return
			case <-timer.C:
			case <-s.openvpn.healthyTimer.C:
				s.onUnhealthyOpenvpn(ctx)
			}
			continue
		}

		// Success, check again in 5 seconds
		const period = 5 * time.Second
		timer := time.NewTimer(period)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return
		case <-timer.C:
		}
	}
}

var (
	errNoIPResolved = errors.New("no IP address resolved")
)

func healthCheck(ctx context.Context, resolver *net.Resolver) (err error) {
	// TODO use mullvad API if current provider is Mullvad
	const domainToResolve = "github.com"
	ips, err := resolver.LookupIP(ctx, "ip", domainToResolve)
	switch {
	case err != nil:
		return err
	case len(ips) == 0:
		return fmt.Errorf("%w for %s", errNoIPResolved, domainToResolve)
	default:
		return nil
	}
}
