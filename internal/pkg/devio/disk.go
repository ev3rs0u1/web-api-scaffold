package devio

import (
	"github.com/shirou/gopsutil/v3/disk"
	"runtime"
	"strings"
	"web-api-scaffold/internal/pkg/constant"
)

type Disk struct {
	Free        uint64   `json:"free"`
	Used        uint64   `json:"used"`
	Total       uint64   `json:"total"`
	UsedPercent float64  `json:"used_percent"`
	MountPoints []string `json:"mount_points"`
}

func FindDiskMountPoints() []string {
	var (
		err    error
		points []string
		parts  []disk.PartitionStat
		usage  *disk.UsageStat
	)

	if parts, err = disk.Partitions(true); err != nil {
		return nil
	}

	for i := range parts {
		if usage, err = disk.Usage(parts[i].Mountpoint); err != nil {
			continue
		}

		if usage.Total >= constant.DeviceSingleDiskSize ||
			runtime.GOOS == "windows" {
			if runtime.GOOS == constant.TaraXSystem {
				if !strings.Contains(usage.Path, "storage") {
					continue
				}
			}

			if runtime.GOOS == constant.TaraWSSystem {
				if !strings.Contains(usage.Path, "smfs") {
					continue
				}
			}

			points = append(points, parts[i].Mountpoint)
		}
	}

	return points
}
