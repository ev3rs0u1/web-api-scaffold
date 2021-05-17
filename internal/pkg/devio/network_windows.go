package devio

func GetMacAddress() string {
	return getMacAddressByInterfaces()
}
