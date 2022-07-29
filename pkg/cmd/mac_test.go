package cmd

import (
	"fmt"
	"testing"
)

func TestMAC(t *testing.T) {
	mac, _ := GetLinuxMACAddress()

	println("mac: ", mac)

	macs := GetMAC()
	fmt.Println("MAC:", macs)

	md5s := MD5Sum(mac)
	fmt.Println("md5s: ", md5s)

	ip, err := GetLocalIP()
	if err != nil {
		return
	}
	fmt.Println("ip: ", ip.String())
}
