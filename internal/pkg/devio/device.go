package devio

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

type DeviceType uint8

const (
	DeviceTaraX  DeviceType = 1
	DeviceTaraM  DeviceType = 2
	DeviceTaraWS DeviceType = 3
)

type Device struct {
	Type      DeviceType `json:"type"`
	BinFile   *BinFile   `json:"binfile"`
	Interface *Interface `json:"interface"`
}

type BinFile struct {
	Version   float64 `json:"version"`
	BuildTag  string  `json:"build_tag"`
	BuildTime int64   `json:"build_time"`
}

type Interface struct {
	LAN  bool `json:"lan"`
	USB  bool `json:"usb"`
	HDMI bool `json:"hdmi"`
}

func GetDeviceType() DeviceType {
	if runtime.GOOS == "android" {
		return DeviceTaraX
	}

	if runtime.GOOS == "linux" &&
		runtime.GOARCH == "arm64" {
		return DeviceTaraM
	}

	return DeviceTaraWS
}

func checkDevIFaceLANOnOff() bool {
	filename := "/sys/class/net/eth0/speed"
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer f.Close()

	var flag byte
	if flag, err = bufio.NewReader(f).ReadByte(); err != nil {
		return false
	}

	return flag != '0'
}

func checkDevIFaceUSBOnOff() bool {
	grep := exec.Command("grep 'type ufsd'")
	mount := exec.Command("mount")

	mount.Stdin, _ = grep.StdoutPipe()

	if err := mount.Start(); err != nil {
		return false
	}

	if err := grep.Run(); err != nil {
		return false
	}

	code := 1
	if err := mount.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				code = status.ExitStatus()
			}
		}
	}

	return code == 0
}

func checkDevIFaceHDMIOnOff() bool {
	filename := "/sys/class/switch/hdmi/state"
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer f.Close()

	var flag byte
	if flag, err = bufio.NewReader(f).ReadByte(); err != nil {
		return false
	}

	return flag == '1'
}
