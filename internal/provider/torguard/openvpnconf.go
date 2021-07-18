package torguard

import (
	"strconv"

	"github.com/rxtreme8/gluetun/internal/configuration"
	"github.com/rxtreme8/gluetun/internal/constants"
	"github.com/rxtreme8/gluetun/internal/models"
	"github.com/rxtreme8/gluetun/internal/provider/utils"
)

func (t *Torguard) BuildConf(connection models.OpenVPNConnection,
	username string, settings configuration.OpenVPN) (lines []string) {
	if settings.Cipher == "" {
		settings.Cipher = constants.AES256gcm
	}

	if settings.Auth == "" {
		settings.Auth = constants.SHA256
	}

	const defaultMSSFix = 1450
	if settings.MSSFix == 0 {
		settings.MSSFix = defaultMSSFix
	}

	lines = []string{
		"client",
		"dev tun",
		"nobind",
		"persist-key",
		"remote-cert-tls server",
		"tls-exit",

		// Torguard specific
		"tun-mtu 1500",
		"tun-mtu-extra 32",
		"mssfix " + strconv.Itoa(int(settings.MSSFix)),
		"reneg-sec 0",
		"key-direction 1",
		"script-security 2",
		"ncp-disable",
		"compress",
		"keepalive 5 30",
		"sndbuf 393216",
		"rcvbuf 393216",

		// Added constant values
		"auth-nocache",
		"mute-replay-warnings",
		"pull-filter ignore \"auth-token\"", // prevent auth failed loops
		"auth-retry nointeract",
		"suppress-timestamps",

		// Modified variables
		"verb " + strconv.Itoa(settings.Verbosity),
		"auth-user-pass " + constants.OpenVPNAuthConf,
		connection.ProtoLine(),
		connection.RemoteLine(),
		"auth " + settings.Auth,
	}

	lines = append(lines, utils.CipherLines(settings.Cipher, settings.Version)...)

	if !settings.Root {
		lines = append(lines, "user "+username)
	}

	if connection.Protocol == constants.UDP {
		lines = append(lines, "fast-io")
	}

	lines = append(lines, utils.WrapOpenvpnCA(
		constants.TorguardCertificate)...)
	lines = append(lines, utils.WrapOpenvpnTLSAuth(
		constants.TorguardOpenvpnStaticKeyV1)...)

	lines = append(lines, "")

	return lines
}
