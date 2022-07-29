package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"net"
	"strings"
)

//获取wlan0的MAC
func GetMAC() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error:" + err.Error())
	}
	var mac string
	for _, inter := range interfaces {
		if strings.EqualFold(inter.Name, "wlan0") {
			mac = inter.HardwareAddr.String()
			break
		}
	}
	return mac
}

// GetLinuxMACAddress 获取mac
func GetLinuxMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err.Error())
	}
	mac, macerr := "", errors.New("无法获取到正确的MAC地址")
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				if ok && ipnet.IP.IsGlobalUnicast() {
					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, macerr
}

// MakeMD5Sum 生成MD5
func MD5Sum(mac string) string {
	ha := md5.New()
	ha.Write([]byte(mac))
	md5s := hex.EncodeToString(ha.Sum(nil))
	md5s = strings.ToUpper(md5s)
	return md5s
}
