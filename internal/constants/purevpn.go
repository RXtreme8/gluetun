package constants

import (
	"net"

	"github.com/rxtreme8/gluetun/internal/models"
)

//nolint:lll
const (
	PurevpnCertificateAuthority = "MIIE6DCCA9CgAwIBAgIJAMjXFoeo5uSlMA0GCSqGSIb3DQEBCwUAMIGoMQswCQYDVQQGEwJISzEQMA4GA1UECBMHQ2VudHJhbDELMAkGA1UEBxMCSEsxGDAWBgNVBAoTD1NlY3VyZS1TZXJ2ZXJDQTELMAkGA1UECxMCSVQxGDAWBgNVBAMTD1NlY3VyZS1TZXJ2ZXJDQTEYMBYGA1UEKRMPU2VjdXJlLVNlcnZlckNBMR8wHQYJKoZIhvcNAQkBFhBtYWlsQGhvc3QuZG9tYWluMB4XDTE2MDExNTE1MzQwOVoXDTI2MDExMjE1MzQwOVowgagxCzAJBgNVBAYTAkhLMRAwDgYDVQQIEwdDZW50cmFsMQswCQYDVQQHEwJISzEYMBYGA1UEChMPU2VjdXJlLVNlcnZlckNBMQswCQYDVQQLEwJJVDEYMBYGA1UEAxMPU2VjdXJlLVNlcnZlckNBMRgwFgYDVQQpEw9TZWN1cmUtU2VydmVyQ0ExHzAdBgkqhkiG9w0BCQEWEG1haWxAaG9zdC5kb21haW4wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDluufhyLlyvXzPUL16kAWAdivl1roQv3QHbuRshyKacf/1Er1JqEbtW3Mx9Fvr/u27qU2W8lQI6DaJhU2BfijPe/KHkib55mvHzIVvoexxya26nk79F2c+d9PnuuMdThWQO3El5a/i2AASnM7T7piIBT2WRZW2i8RbfJaTT7G7LP7OpMKIV1qyBg/cWoO7cIWQW4jmzqrNryIkF0AzStLN1DxvnQZwgXBGv0CwuAkfQuNSLu0PQgPp0PhdukNZFllv5D29IhPr0Z+kwPtrAgPQo+lHlOBHBMUpDT4XChTPeAvMaUSBsqmonAE8UUHEabWrqYN/kWNHCNkYXMkiVmK1AgMBAAGjggERMIIBDTAdBgNVHQ4EFgQU456ijsFrYnzHBShLAPpOUqQ+Z2cwgd0GA1UdIwSB1TCB0oAU456ijsFrYnzHBShLAPpOUqQ+Z2ehga6kgaswgagxCzAJBgNVBAYTAkhLMRAwDgYDVQQIEwdDZW50cmFsMQswCQYDVQQHEwJISzEYMBYGA1UEChMPU2VjdXJlLVNlcnZlckNBMQswCQYDVQQLEwJJVDEYMBYGA1UEAxMPU2VjdXJlLVNlcnZlckNBMRgwFgYDVQQpEw9TZWN1cmUtU2VydmVyQ0ExHzAdBgkqhkiG9w0BCQEWEG1haWxAaG9zdC5kb21haW6CCQDI1xaHqObkpTAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCvga2HMwOtUxWH/inL2qk24KX2pxLg939JNhqoyNrUpbDHag5xPQYXUmUpKrNJZ0z+o/ZnNUPHydTSXE7Z7E45J0GDN5E7g4pakndKnDLSjp03NgGsCGW+cXnz6UBPM5FStFvGdDeModeSUyoS9fjk+mYROvmiy5EiVDP91sKGcPLR7Ym0M7zl2aaqV7bb98HmMoBOxpeZQinof67nKrCsgz/xjktWFgcmPl4/PQSsmqQD0fTtWxGuRX+FzwvF2OCMCAJgp1RqJNlk2g50/kBIoJVPPCfjDFeDU5zGaWGSQ9+z1L6/z7VXdjUiHL0ouOcHwbiS4ZjTr9nMn6WdAHU2"
	PurevpnCertificate          = "MIIEnzCCA4egAwIBAgIBAzANBgkqhkiG9w0BAQsFADCBqDELMAkGA1UEBhMCSEsxEDAOBgNVBAgTB0NlbnRyYWwxCzAJBgNVBAcTAkhLMRgwFgYDVQQKEw9TZWN1cmUtU2VydmVyQ0ExCzAJBgNVBAsTAklUMRgwFgYDVQQDEw9TZWN1cmUtU2VydmVyQ0ExGDAWBgNVBCkTD1NlY3VyZS1TZXJ2ZXJDQTEfMB0GCSqGSIb3DQEJARYQbWFpbEBob3N0LmRvbWFpbjAeFw0xNjAxMTUxNjE1MzhaFw0yNjAxMTIxNjE1MzhaMIGdMQswCQYDVQQGEwJISzEQMA4GA1UECBMHQ2VudHJhbDELMAkGA1UEBxMCSEsxFjAUBgNVBAoTDVNlY3VyZS1DbGllbnQxCzAJBgNVBAsTAklUMRYwFAYDVQQDEw1TZWN1cmUtQ2xpZW50MREwDwYDVQQpEwhjaGFuZ2VtZTEfMB0GCSqGSIb3DQEJARYQbWFpbEBob3N0LmRvbWFpbjCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAxsnyn4v6xxDPnuDaYS0b9M1N8nxgg7OBPBlK+FWRxdTQ8yxt5U5CZGm7riVp7fya2J2iPZIgmHQEv/KbxztsHAVlYSfYYlalrnhEL3bDP2tY+N43AwB1k5BrPq2s1pPLT2XG951drDKG4PUuFHUP1sHzW5oQlfVCmxgIMAP8OYkCAwEAAaOCAV8wggFbMAkGA1UdEwQCMAAwLQYJYIZIAYb4QgENBCAWHkVhc3ktUlNBIEdlbmVyYXRlZCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQU9MwUnUDbQKKZKjoeieD2OD5NlAEwgd0GA1UdIwSB1TCB0oAU456ijsFrYnzHBShLAPpOUqQ+Z2ehga6kgaswgagxCzAJBgNVBAYTAkhLMRAwDgYDVQQIEwdDZW50cmFsMQswCQYDVQQHEwJISzEYMBYGA1UEChMPU2VjdXJlLVNlcnZlckNBMQswCQYDVQQLEwJJVDEYMBYGA1UEAxMPU2VjdXJlLVNlcnZlckNBMRgwFgYDVQQpEw9TZWN1cmUtU2VydmVyQ0ExHzAdBgkqhkiG9w0BCQEWEG1haWxAaG9zdC5kb21haW6CCQDI1xaHqObkpTATBgNVHSUEDDAKBggrBgEFBQcDAjALBgNVHQ8EBAMCB4AwDQYJKoZIhvcNAQELBQADggEBAFyFo2VUX/UFixsdPdK9/Yt6mkCWc+XS1xbapGXXb9U1d+h1iBCIV9odUHgNCXWpz1hR5Uu/OCzaZ0asLE4IFMZlQmJs8sMT0c1tfPPGW45vxbL0lhqnQ8PNcBH7huNK7VFjUh4szXRKmaQPaM4S91R3L4CaNfVeHfAg7mN2m9Zn5Gto1Q1/CFMGKu2hxwGEw5p+X1czBWEvg/O09ckx/ggkkI1NcZsNiYQ+6Pz8DdGGX3+05YwLZu94+O6iIMrzxl/il0eK83g3YPbsOrASARvw6w/8sOnJCK5eOacl21oww875KisnYdWjHB1FiI+VzQ1/gyoDsL5kPTJVuu2CoG8="
	PurevpnKey                  = "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMbJ8p+L+scQz57g2mEtG/TNTfJ8YIOzgTwZSvhVkcXU0PMsbeVOQmRpu64lae38mtidoj2SIJh0BL/ym8c7bBwFZWEn2GJWpa54RC92wz9rWPjeNwMAdZOQaz6trNaTy09lxvedXawyhuD1LhR1D9bB81uaEJX1QpsYCDAD/DmJAgMBAAECgYEAvTHbDupE5U0krUvHzBEIuHblptGlcfNYHoDcD3oxYR3pOGeiuElBexv+mgHVzcFLBrsQfJUlHLPfCWi3xmjRvDQcr7N7U1u7NIzazy/PpRBaKolMRiM1KMYi2DG0i4ZONwFT8bvNHOIrZzCLY54KDrqOn55OzC70WYjWh4t5evkCQQDkkzZUAeskBC9+JP/zLps8jhwfoLBWGw/zbC9ePDmX0N8MTZdcUpg6KUTf1wbkLUyVtIRjS2ao6qu1jWG6K0x3AkEA3qPWyaWQWCynhNDqu2U1cPb2kh5AJip+gqxO3emikAdajsSxeoyEC2AfyBITbeB1tvCUZH17J4i/0+OFTEQp/wJAb/zEOGJ8PzghwK8GC7JA8mk51DEZVAaMSRovFv9wxDXcoh191AjPdmdzzCuAv9iF1i8MUc3GbWoUWK39PIYsPwJAWh63sqfx5b8tj/WBDpnJKBDPfhYAoXJSA1L8GZeY1fQkE+ZKcPCwAmrGcpXeh3t0Krj3WDXyw+32uC5Apr5wwQJAPZwOOReaC4YNfBPZN9BdHvVjOYGGUffpI+X+hWpLRnQFJteAi+eqwyk0Oi0SkJB+a7jcerK2d7q7xhec5WHlng=="
	PurevpnOpenvpnStaticKeyV1   = "e30af995f56d07426d9ba1f824730521d4283db4b4d0cdda9c6e8759a3799dcb7939b6a5989160c9660de0f6125cbb1f585b41c074b2fe88ecfcf17eab9a33be1352379cdf74952b588fb161a93e13df9135b2b29038231e02d657a6225705e6868ccb0c384ed11614690a1894bfbeb274cebf1fe9c2329bdd5c8a40fe8820624d2ea7540cd79ab76892db51fc371a3ac5fc9573afecb3fffe3281e61d72e91579d9b03d8cbf7909b3aebf4d90850321ee6b7d0a7846d15c27d8290e031e951e19438a4654663cad975e138f5bc5af89c737ad822f27e19057731f41e1e254cc9c95b7175c622422cde9f1f2cfd3510add94498b4d7133d3729dd214a16b27fb"
)

func PurevpnRegionChoices() (choices []string) {
	servers := PurevpnServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return makeUnique(choices)
}

func PurevpnCountryChoices() (choices []string) {
	servers := PurevpnServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func PurevpnCityChoices() (choices []string) {
	servers := PurevpnServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return makeUnique(choices)
}

func PurevpnHostnameChoices() (choices []string) {
	servers := PurevpnServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}

//nolint:lll
// PurevpnServers returns a slice of all the server information for Purevpn.
func PurevpnServers() []models.PurevpnServer {
	return []models.PurevpnServer{
		{Country: "Australia", Region: "New South Wales", City: "Sydney", Hostname: "au-sd2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{43, 245, 163, 82}, {46, 243, 245, 4}}},
		{Country: "Australia", Region: "New South Wales", City: "Sydney", Hostname: "au-sd2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{43, 245, 161, 80}, {43, 245, 161, 81}, {43, 245, 161, 82}, {43, 245, 161, 86}}},
		{Country: "Australia", Region: "Western Australia", City: "Perth", Hostname: "au2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{43, 250, 205, 50}}},
		{Country: "Australia", Region: "Western Australia", City: "Perth", Hostname: "au2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{43, 250, 205, 50}, {43, 250, 205, 61}}},
		{Country: "Australia", Region: "Western Australia", City: "Perth", Hostname: "au2-pe-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{43, 250, 205, 50}, {43, 250, 205, 51}}},
		{Country: "Australia", Region: "Western Australia", City: "Perth", Hostname: "au2-pe-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{43, 250, 205, 34}}},
		{Country: "Austria", Region: "Vienna", City: "Vienna", Hostname: "at2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{217, 64, 127, 251}, {217, 64, 127, 252}}},
		{Country: "Austria", Region: "Vienna", City: "Vienna", Hostname: "at2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{37, 120, 212, 219}, {217, 64, 127, 252}}},
		{Country: "Belgium", Region: "Flanders", City: "Zaventem", Hostname: "vleu-be2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{217, 138, 211, 87}, {217, 138, 211, 116}, {217, 138, 221, 117}}},
		{Country: "Brazil", Region: "São Paulo", City: "São Paulo", Hostname: "br2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{181, 41, 201, 83}, {181, 41, 201, 84}}},
		{Country: "Brazil", Region: "São Paulo", City: "São Paulo", Hostname: "br2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{138, 99, 210, 4}, {181, 41, 201, 75}, {181, 41, 201, 82}, {181, 41, 201, 83}}},
		{Country: "Bulgaria", Region: "Sofia-Capital", City: "Sofia", Hostname: "vleu-be2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{217, 138, 221, 114}}},
		{Country: "Canada", Region: "British Columbia", City: "Vancouver", Hostname: "cav2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{66, 115, 147, 34}}},
		{Country: "Canada", Region: "British Columbia", City: "Vancouver", Hostname: "cav2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{66, 115, 147, 34}, {66, 115, 147, 35}, {66, 115, 147, 37}, {66, 115, 147, 38}, {66, 115, 147, 40}, {66, 115, 147, 43}}},
		{Country: "Canada", Region: "Ontario", City: "Toronto", Hostname: "ca2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{104, 200, 138, 212}}},
		{Country: "Canada", Region: "Ontario", City: "Toronto", Hostname: "ca2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{104, 200, 138, 213}, {104, 200, 138, 214}, {104, 200, 138, 216}}},
		{Country: "Czech Republic", Region: "Hlavní město Praha", City: "Prague", Hostname: "cz2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{185, 156, 174, 36}}},
		{Country: "Czech Republic", Region: "Hlavní město Praha", City: "Prague", Hostname: "cz2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{185, 156, 174, 38}, {217, 138, 199, 252}}},
		{Country: "France", Region: "Île-de-France", City: "Paris", Hostname: "fr2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{45, 152, 181, 68}}},
		{Country: "France", Region: "Île-de-France", City: "Paris", Hostname: "fr2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{45, 152, 181, 67}, {45, 152, 181, 70}}},
		{Country: "Germany", Region: "Hesse", City: "Frankfurt am Main", Hostname: "de2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{5, 254, 88, 213}}},
		{Country: "Germany", Region: "Hesse", City: "Frankfurt am Main", Hostname: "de2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 254, 82, 50}, {5, 254, 82, 52}}},
		{Country: "Hong Kong", Region: "Central and Western", City: "Hong Kong", Hostname: "hk2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{103, 109, 103, 59}, {103, 109, 103, 60}, {119, 81, 228, 109}}},
		{Country: "Hong Kong", Region: "Central and Western", City: "Hong Kong", Hostname: "hk2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 111, 168, 4}}},
		{Country: "Hong Kong", Region: "Central and Western", City: "Hong Kong", Hostname: "vlap-ph2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{36, 255, 97, 3}, {129, 227, 119, 84}}},
		{Country: "Hong Kong", Region: "Central and Western", City: "Hong Kong", Hostname: "vlap-ph2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{36, 255, 97, 3}}},
		{Country: "Italy", Region: "Lombardy", City: "Milan", Hostname: "it2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{172, 111, 173, 3}}},
		{Country: "Italy", Region: "Lombardy", City: "Milan", Hostname: "it2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 111, 173, 3}}},
		{Country: "Japan", Region: "Tokyo", City: "Tokyo", Hostname: "jp2-ovpn.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{185, 242, 4, 59}}},
		{Country: "Korea", Region: "Gangwon-do", City: "Hongch’ŏn", Hostname: "kr2-ovpn.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{43, 226, 231, 4}, {43, 226, 231, 6}}},
		{Country: "Malaysia", Region: "Kuala Lumpur", City: "Kuala Lumpur", Hostname: "my2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{103, 55, 10, 134}}},
		{Country: "Malaysia", Region: "Kuala Lumpur", City: "Kuala Lumpur", Hostname: "my2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{103, 28, 90, 32}}},
		{Country: "Netherlands", Region: "North Holland", City: "Amsterdam", Hostname: "nl2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{5, 254, 73, 172}, {5, 254, 73, 252}, {92, 119, 179, 195}, {206, 123, 147, 4}}},
		{Country: "Netherlands", Region: "North Holland", City: "Amsterdam", Hostname: "nl2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 254, 73, 171}, {5, 254, 73, 202}, {92, 119, 179, 195}, {206, 123, 147, 4}}},
		{Country: "Netherlands", Region: "North Holland", City: "Amsterdam", Hostname: "vlap-th2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{104, 37, 6, 4}}},
		{Country: "Netherlands", Region: "North Holland", City: "Amsterdam", Hostname: "vlap-th2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{104, 37, 6, 4}}},
		{Country: "Norway", Region: "Oslo", City: "Oslo", Hostname: "vleu-no2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{82, 102, 22, 212}}},
		{Country: "Norway", Region: "Oslo", City: "Oslo", Hostname: "vleu-no2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{82, 102, 22, 212}}},
		{Country: "Poland", Region: "Mazovia", City: "Warsaw", Hostname: "pl2-ovpn.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 253, 206, 251}}},
		{Country: "Portugal", Region: "Lisbon", City: "Lisbon", Hostname: "Pt2-ovpn.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 154, 174, 3}}},
		{Country: "Russian Federation", Region: "St.-Petersburg", City: "Saint Petersburg", Hostname: "ru2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{94, 242, 54, 23}, {206, 123, 139, 4}}},
		{Country: "Russian Federation", Region: "St.-Petersburg", City: "Saint Petersburg", Hostname: "ru2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{94, 242, 54, 23}, {206, 123, 139, 4}}},
		{Country: "Singapore", Region: "Singapore", City: "Singapore", Hostname: "sg2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{84, 247, 49, 181}}},
		{Country: "Singapore", Region: "Singapore", City: "Singapore", Hostname: "sg2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{37, 120, 208, 147}, {37, 120, 208, 148}}},
		{Country: "South Africa", Region: "Gauteng", City: "Johannesburg", Hostname: "za2-ovpn.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{102, 165, 3, 33}, {102, 165, 3, 34}}},
		{Country: "Sweden", Region: "Stockholm", City: "Stockholm", Hostname: "se2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{45, 12, 220, 123}, {86, 106, 103, 178}, {86, 106, 103, 179}}},
		{Country: "Sweden", Region: "Stockholm", City: "Stockholm", Hostname: "se2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{45, 12, 220, 125}, {86, 106, 103, 183}}},
		{Country: "Switzerland", Region: "Zurich", City: "Zürich", Hostname: "ch2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{45, 12, 222, 98}, {45, 12, 222, 99}}},
		{Country: "Switzerland", Region: "Zurich", City: "Zürich", Hostname: "ch2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{45, 12, 222, 100}, {45, 12, 222, 102}, {45, 12, 222, 106}, {45, 12, 222, 107}}},
		{Country: "Taiwan", Region: "Taiwan", City: "Taipei", Hostname: "tw2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{128, 1, 155, 178}}},
		{Country: "Taiwan", Region: "Taiwan", City: "Taipei", Hostname: "tw2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{128, 1, 155, 178}}},
		{Country: "United Kingdom", Region: "England", City: "London", Hostname: "ukg2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{45, 74, 0, 4}}},
		{Country: "United Kingdom", Region: "England", City: "London", Hostname: "ukl2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{45, 74, 0, 4}}},
		{Country: "United Kingdom", Region: "England", City: "Manchester", Hostname: "ukg2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{89, 238, 165, 148}, {89, 238, 165, 149}}},
		{Country: "United Kingdom", Region: "England", City: "Manchester", Hostname: "ukl2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{89, 238, 165, 148}, {89, 238, 165, 149}}},
		{Country: "United Kingdom", Region: "England", City: "Manchester", Hostname: "ukm2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{188, 72, 89, 4}}},
		{Country: "United Kingdom", Region: "England", City: "Manchester", Hostname: "ukm2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{188, 72, 89, 4}}},
		{Country: "United States", Region: "California", City: "Los Angeles", Hostname: "usla2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{89, 45, 4, 2}}},
		{Country: "United States", Region: "California", City: "Los Angeles", Hostname: "vlus-mx2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{104, 243, 243, 131}}},
		{Country: "United States", Region: "California", City: "Los Angeles", Hostname: "vlus-mx2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{104, 243, 243, 131}}},
		{Country: "United States", Region: "California", City: "San Jose", Hostname: "ussf2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{141, 101, 166, 4}}},
		{Country: "United States", Region: "Florida", City: "Miami", Hostname: "usfl2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{5, 254, 79, 100}, {5, 254, 79, 116}}},
		{Country: "United States", Region: "Florida", City: "Miami", Hostname: "usfl2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 254, 79, 114}, {5, 254, 79, 115}}},
		{Country: "United States", Region: "Georgia", City: "Cumming", Hostname: "in2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{169, 38, 69, 12}}},
		{Country: "United States", Region: "New Jersey", City: "Piscataway", Hostname: "usny2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{209, 127, 184, 71}, {209, 127, 184, 73}, {209, 127, 184, 74}, {209, 127, 184, 76}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "usnj2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{172, 94, 72, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "usnj2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 94, 72, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "usny2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{141, 101, 153, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "ussf2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 94, 72, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "uswdc2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 94, 72, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "vlus-pa2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{172, 111, 149, 4}}},
		{Country: "United States", Region: "New York", City: "New York City", Hostname: "vlus-pa2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 111, 149, 4}}},
		{Country: "United States", Region: "Pennsylvania", City: "Pittsburgh", Hostname: "in2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{104, 250, 160, 132}}},
		{Country: "United States", Region: "Texas", City: "Allen", Hostname: "usga2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{172, 94, 108, 4}}},
		{Country: "United States", Region: "Texas", City: "Allen", Hostname: "usil2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{172, 94, 108, 4}}},
		{Country: "United States", Region: "Texas", City: "Dallas", Hostname: "ustx2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{208, 84, 155, 100}, {208, 84, 155, 101}}},
		{Country: "United States", Region: "Texas", City: "Dallas", Hostname: "ustx2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{104, 217, 255, 185}, {104, 217, 255, 186}}},
		{Country: "United States", Region: "Virginia", City: "Centreville", Hostname: "usva2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{5, 254, 109, 115}, {37, 221, 173, 212}}},
		{Country: "United States", Region: "Virginia", City: "Reston", Hostname: "usva2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{5, 254, 77, 26}, {5, 254, 77, 27}}},
		{Country: "United States", Region: "Washington", City: "Seattle", Hostname: "ussa2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{172, 94, 86, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "us2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{46, 243, 249, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "us2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}, {172, 94, 72, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usga2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usil2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}, {172, 94, 72, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usla2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}, {172, 94, 72, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usphx2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "ussa2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{46, 243, 249, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usut2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{46, 243, 249, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "usut2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{141, 101, 169, 4}, {172, 94, 72, 4}, {172, 94, 86, 4}}},
		{Country: "United States", Region: "Washington, D.C.", City: "Washington", Hostname: "uswdc2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{46, 243, 249, 4}}},
		{Country: "Vietnam", Region: "Hanoi", City: "Hanoi", Hostname: "vlap-vn2-ovpn-tcp.pointtoserver.com", TCP: true, UDP: false, IPs: []net.IP{{192, 253, 249, 132}}},
		{Country: "Vietnam", Region: "Hanoi", City: "Hanoi", Hostname: "vlap-vn2-ovpn-udp.pointtoserver.com", TCP: false, UDP: true, IPs: []net.IP{{192, 253, 249, 132}}},
	}
}
