package devio

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"net"
)

type Network struct {
	LanIP   string `json:"lan_ip"`
	WanIP   string `json:"wan_ip"`
	MacAddr string `json:"mac_addr"`
}

func GetExternalIP() string {
	conn, err := net.DialTimeout("tcp",
		"ns1.dnspod.net:6666", TCPTimeout)
	if err != nil {
		_, body, errs := fiber.Get("http://pv.sohu.com/cityjson?ie=utf-8").Bytes()
		if len(errs) > 0 {
			return ""
		}

		items := bytes.Split(body, []byte(`"`))
		for i := range items {
			if net.ParseIP(string(items[i])) != nil {
				body = bytes.TrimSpace(items[i])
				break
			}
		}

		return string(body)
	}
	defer conn.Close()

	addr := make([]byte, 32)
	n, err := conn.Read(addr)
	if err != nil {
		return ""
	}

	return string(bytes.TrimSpace(addr[:n]))
}

func GetInternalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range interfaces {
		adders, err := i.Addrs()
		if err != nil {
			return ""
		}

		for _, addr := range adders {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				return v.String()
			}

			if isPrivateIP(ip) {
				return ip.String()
			}
		}
	}

	return ""
}

func isPrivateIP(ip net.IP) bool {
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func getMacAddressByInterfaces() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range interfaces {
		adders, err := i.Addrs()
		if err != nil {
			return ""
		}

		for _, addr := range adders {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if isPrivateIP(ip) {
				return i.HardwareAddr.String()
			}
		}
	}

	return ""
}
