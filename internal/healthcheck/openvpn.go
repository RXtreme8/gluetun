package healthcheck

import (
	"context"
	"time"

	"github.com/rxtreme8/gluetun/internal/constants"
)

func (s *server) onUnhealthyOpenvpn(ctx context.Context) {
	s.logger.Info("program has been unhealthy for " +
		s.openvpn.healthyWaitTime.String() + ": restarting OpenVPN")
	_, _ = s.openvpn.looper.ApplyStatus(ctx, constants.Stopped)
	_, _ = s.openvpn.looper.ApplyStatus(ctx, constants.Running)
	s.openvpn.healthyWaitTime += openvpnHealthyWaitTimeAdd
	s.openvpn.healthyTimer = time.NewTimer(s.openvpn.healthyWaitTime)
}
