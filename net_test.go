package alphago

import (
	"reflect"
	"testing"
)

var (
	startIp = "255.129.34.8"
	endIp   = "255.129.34.8"
	longIp  = uint32(4286652936)
)

func Test_Ip2Long_1(t *testing.T) {
	if 4286652936 != Ip2Long(startIp) {
		t.Error(`Ip2Long(255.129.34.8) != 4286652936`)
	}
}

func Test_Ip2Long_2(t *testing.T) {
	if 0 != Ip2Long("99999.9.0") {
		t.Error(`Ip2Long("99999.9.0") != 0`)
	}
}

func Test_Ip2Long_3(t *testing.T) {
	if 0 != Ip2Long("89:89:90:9") {
		t.Error(`Ip2Long("89:89:90:9") != 0`)
	}
}

func Test_Long2Ip(t *testing.T) {
	if "255.129.34.8" != Long2Ip(longIp) {
		t.Error(`Long2Ip(uint32(4286652936)) != "255.129.34.8"`)
	}
}

func Test_IpRange_1(t *testing.T) {
	var ips = []string{"10.34.56.78", "10.34.56.79", "10.34.56.80"}
	if !reflect.DeepEqual(IpRange("10.34.56.78", "10.34.56.80"), ips) {
		t.Error("IpRange error")
	}
}

func Test_IpRange_2(t *testing.T) {
	if 0 != len(IpRange("103.456.78", "10.34.56.90")) {
		t.Error(`IpRange("103.456.78", "10.34.56.90") != []`)
	}
}

func Test_IpRange_3(t *testing.T) {
	if 0 != len(IpRange("10.34.56.78", "10.34.56:90")) {
		t.Error(`IpRange("10.34.56.78", "10.34.56:90") != []`)
	}
}

func Test_IpRange_4(t *testing.T) {
	if 0 != len(IpRange("10.34.56:78", "10.34.56.80")) {
		t.Error(`IpRange("10.34.56:78", "10.34.56.80") != []`)
	}
}

func Test_IsPrivate_1(t *testing.T) {
	if true != IsPrivate("10.0.0.0") {
		t.Error(`IsPrivate("10.0.0.0") is not true`)
	}
}

func Test_IsPrivate_2(t *testing.T) {
	if true != IsPrivate("10.0.0.255") {
		t.Error(`IsPrivate("10.0.0.0") is not true`)
	}
}

func Test_IsPrivate_3(t *testing.T) {
	if false != IsPrivate("8.8.8.8") {
		t.Error(`IsPrivate("8.8.8.8") is not false`)
	}
}

func Test_IsPrivate_4(t *testing.T) {
	if true != IsPrivate("172.16.0.1") {
		t.Error(`IsPrivate("172.16.0.1") is not true`)
	}
}

func Test_IsPrivate_5(t *testing.T) {
	if true != IsPrivate("192.168.0.1") {
		t.Error(`IsPrivate("192.168.0.1") is not true`)
	}
}

func Test_IsPrivate_6(t *testing.T) {
	if false != IsPrivate("192168.0.1") {
		t.Error(`IsPrivate("192168.0.1") is not false`)
	}
}
