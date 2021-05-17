package devio

import (
	"time"
)

type System struct {
	OS           string `json:"os"`
	BootTime     uint64 `json:"boot_time"`
	SerialNumber string `json:"serial_number"`
}

func GetUnixTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
