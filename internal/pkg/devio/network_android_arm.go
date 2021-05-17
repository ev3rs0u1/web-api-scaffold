package devio

import (
	"bytes"
	"os"
)

func GetMacAddress() string {
	filename := "/sys/class/net/wlan0/address"

	f, err := os.Open(filename)
	if err != nil {
		return getMacAddressByInterfaces()
	}
	defer f.Close()

	addr := make([]byte, 48)
	n, err := f.Read(addr)
	if err != nil {
		return ""
	}

	return string(bytes.TrimSpace(addr[:n]))
}
