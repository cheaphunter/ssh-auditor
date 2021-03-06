package sshauditor

import (
	"net"
	"strings"
	"time"
)

type ScanResult struct {
	hostport string
	success  bool
	banner   string
}

func ScanPort(hostport string) ScanResult {
	res := ScanResult{hostport: hostport}
	var banner string
	conn, err := net.DialTimeout("tcp", hostport, 2*time.Second)
	if err != nil {
		return res
	}
	defer conn.Close()
	bannerBuffer := make([]byte, 256)
	conn.SetDeadline(time.Now().Add(4 * time.Second))
	n, err := conn.Read(bannerBuffer)
	if err == nil {
		banner = string(bannerBuffer[:n])

		newlinePosition := strings.IndexAny(banner, "\r\n")
		if newlinePosition != -1 {
			banner = banner[:newlinePosition]
		}
	}
	res.success = true
	res.banner = banner
	return res
}
