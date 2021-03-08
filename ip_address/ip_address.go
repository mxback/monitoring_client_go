package ip_address

import (
	"net"
)

type IpAddress struct {
	IP string
}

func GetOutboundIP() (ip *IpAddress, err error) {
	// 获取出口地址
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = &IpAddress{
		IP: localAddr.IP.String(),
	}
	return
}
