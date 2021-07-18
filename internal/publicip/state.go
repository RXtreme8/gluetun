package publicip

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"sync"

	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/models"
)

type state struct {
	status     models.LoopStatus
	settings   configuration.PublicIP
	ip         net.IP
	statusMu   sync.RWMutex
	settingsMu sync.RWMutex
	ipMu       sync.RWMutex
}

func (s *state) setStatusWithLock(status models.LoopStatus) {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()
	s.status = status
}

func (l *looper) GetStatus() (status models.LoopStatus) {
	l.state.statusMu.RLock()
	defer l.state.statusMu.RUnlock()
	return l.state.status
}

var ErrInvalidStatus = errors.New("invalid status")

func (l *looper) SetStatus(ctx context.Context, status models.LoopStatus) (
	outcome string, err error) {
	l.state.statusMu.Lock()
	defer l.state.statusMu.Unlock()
	existingStatus := l.state.status

	switch status {
	case constants.Running:
		switch existingStatus {
		case constants.Starting, constants.Running, constants.Stopping, constants.Crashed:
			return fmt.Sprintf("already %s", existingStatus), nil
		}
		l.loopLock.Lock()
		defer l.loopLock.Unlock()
		l.state.status = constants.Starting
		l.state.statusMu.Unlock()
		l.start <- struct{}{}

		newStatus := constants.Starting // for canceled context
		select {
		case <-ctx.Done():
		case newStatus = <-l.running:
		}
		l.state.statusMu.Lock()
		l.state.status = newStatus
		return newStatus.String(), nil
	case constants.Stopped:
		switch existingStatus {
		case constants.Stopped, constants.Stopping, constants.Starting, constants.Crashed:
			return fmt.Sprintf("already %s", existingStatus), nil
		}
		l.loopLock.Lock()
		defer l.loopLock.Unlock()
		l.state.status = constants.Stopping
		l.state.statusMu.Unlock()
		l.stop <- struct{}{}

		newStatus := constants.Stopping // for canceled context
		select {
		case <-ctx.Done():
		case <-l.stopped:
			newStatus = constants.Stopped
		}
		l.state.statusMu.Lock()
		l.state.status = newStatus
		return status.String(), nil
	default:
		return "", fmt.Errorf("%w: %s: it can only be one of: %s, %s",
			ErrInvalidStatus, status, constants.Running, constants.Stopped)
	}
}

func (l *looper) GetSettings() (settings configuration.PublicIP) {
	l.state.settingsMu.RLock()
	defer l.state.settingsMu.RUnlock()
	return l.state.settings
}

func (l *looper) SetSettings(settings configuration.PublicIP) (
	outcome string) {
	l.state.settingsMu.Lock()
	defer l.state.settingsMu.Unlock()
	settingsUnchanged := reflect.DeepEqual(settings, l.state.settings)
	if settingsUnchanged {
		return "settings left unchanged"
	}
	periodChanged := l.state.settings.Period != settings.Period
	l.state.settings = settings
	if periodChanged {
		l.updateTicker <- struct{}{}
		// TODO blocking
	}
	return "settings updated"
}

func (l *looper) GetPublicIP() (publicIP net.IP) {
	l.state.ipMu.RLock()
	defer l.state.ipMu.RUnlock()
	publicIP = make(net.IP, len(l.state.ip))
	copy(publicIP, l.state.ip)
	return publicIP
}

func (s *state) setPublicIP(publicIP net.IP) {
	s.ipMu.Lock()
	defer s.ipMu.Unlock()
	s.ip = make(net.IP, len(publicIP))
	copy(s.ip, publicIP)
}
