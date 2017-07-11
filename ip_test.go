package alphago

import (
	"testing"
)

var (
	startIp = "255.129.34.8"
	endIp   = "255.129.34.8"
	longIp  = uint32(4286652936)
)

func Test_Ip2Long_1(t *testing.T) {
	_ = Ip2Long(startIp)
}

func Test_Ip2Long_2(t *testing.T) {
	_ = Ip2Long("99999.9.0")
}

func Test_Ip2Long_3(t *testing.T) {
	_ = Ip2Long("89:89:90:9")
}

func Test_Long2Ip(t *testing.T) {
	_ = Long2Ip(longIp)
}

func Test_IpRange_1(t *testing.T) {
	_ = IpRange("10.34.56.78", "10.34.56.90")
}

func Test_IpRange_2(t *testing.T) {
	_ = IpRange("103.456.78", "10.34.56.90")
}

func Test_IpRange_3(t *testing.T) {
	_ = IpRange("10.34.56.78", "10.34.56:90")
}

func Test_IpRange_4(t *testing.T) {
	_ = IpRange("10.34.56:78", "10.34.56.90")
}

func Test_IsPrivate_1(t *testing.T) {
	_ = IsPrivate("10.0.0.0")
}

func Test_IsPrivate_2(t *testing.T) {
	_ = IsPrivate("10.0.0.255")
}

func Test_IsPrivate_3(t *testing.T) {
	_ = IsPrivate("8.8.8.8")
}

func Test_IsPrivate_4(t *testing.T) {
	_ = IsPrivate("172.16.0.1")
}

func Test_IsPrivate_5(t *testing.T) {
	_ = IsPrivate("192.168.0.1")
}

func Test_IsPrivate_6(t *testing.T) {
	_ = IsPrivate("192168.0.1")
}
