package alphago

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

// Ip2Long is a function that converts ip(8.8.8.8 or 255.255.234.23)
// into a long integer.
func Ip2Long(ipString string) (ip uint32) {
	ipParsed := net.ParseIP(ipString)
	if ipParsed == nil {
		return 0
	}
	ipParsed = ipParsed.To4()
	return binary.BigEndian.Uint32(ipParsed)
}

// Long2Ip is a function that converts long integer(4286652936)
// into ip.
func Long2Ip(ipLong uint32) (ip string) {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ipNet := net.IP(ipByte)
	return ipNet.String()
}

// IpRange is a function that calculates all the Ip between
// the start and end ip.And it returns []string{Ips}.
func IpRange(startIp, endIp string) (ips []string) {
	startIpLong := Ip2Long(startIp)
	if startIpLong == 0 && startIp != "0.0.0.0" {
		return ips
	}
	endIpLong := Ip2Long(endIp)
	if endIpLong == 0 && endIp != "0.0.0.0" {
		return ips
	}
	for iterater := startIpLong; iterater <= endIpLong; iterater++ {
		ips = append(ips, Long2Ip(iterater))
	}
	return ips
}

// IsPrivate is a function that determines whether a given ip is a private ip.
func IsPrivate(ipString string) bool {
	ipParsed := net.ParseIP(ipString)
	if ipParsed == nil {
		return false
	}
	ipSplits := strings.Split(ipParsed.String(), ".")
	ipValue2, _ := strconv.Atoi(ipSplits[1])

	switch {
	case ipSplits[0] == "10": // 10.0.0.0/8
		return true
	case ipSplits[0] == "172" && (ipValue2 >= 16 && ipValue2 < 32): // 172.16.0.0/12
		return true
	case ipSplits[0] == "192" && ipValue2 == 168: // 192.168.0.0/16
		return true
	}

	return false
}
