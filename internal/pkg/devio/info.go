package devio

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shopspring/decimal"
	"os"
	"runtime"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/hasher"
)

type Info struct {
	*Device  `json:"device"`
	*Disk    `json:"disk"`
	*System  `json:"system"`
	*Network `json:"network"`
}

func Information() interface{} {
	var (
		points  []string
		usage   *disk.UsageStat
		err     error
		free    uint64
		used    uint64
		total   uint64
		percent float64
	)

	points = FindDiskMountPoints()

	for i := range points {
		if usage, err = disk.Usage(points[i]); err != nil {
			continue
		}

		used += usage.Used
		total += usage.Total
	}

	free = total - used
	percent, _ = decimal.NewFromInt(int64(used)).
		Div(decimal.NewFromInt(int64(total))).
		Shift(2).
		Truncate(2).
		Float64()

	bootTime, _ := host.BootTime()
	macAddr := GetMacAddress()

	exec, _ := os.Executable()
	file, _ := os.Open(exec)
	fi, _ := file.Stat()
	defer file.Close()

	var info = &Info{
		Device: &Device{
			Type: GetDeviceType(),
			BinFile: &BinFile{
				Version:   constant.BinFileVersion,
				BuildTag:  hasher.CalculateSmallHashByReader(file),
				BuildTime: fi.ModTime().Unix(),
			},
			Interface: &Interface{
				LAN:  checkDevIFaceLANOnOff(),
				USB:  checkDevIFaceUSBOnOff(),
				HDMI: checkDevIFaceHDMIOnOff(),
			},
		},
		Disk: &Disk{
			Free:        free,
			Used:        used,
			Total:       total,
			UsedPercent: percent,
			MountPoints: points,
		},
		System: &System{
			OS:           runtime.GOOS,
			BootTime:     bootTime,
			SerialNumber: hasher.CalculateMD5Hash(macAddr),
		},
		Network: &Network{
			WanIP:   GetExternalIP(),
			LanIP:   GetInternalIP(),
			MacAddr: macAddr,
		},
	}

	return info
}
