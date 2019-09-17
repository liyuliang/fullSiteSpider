package format

import "testing"

func TestIpToInt64(t *testing.T) {
	ip := "192.168.78.123"

	ipInt, err := IpToInt64(ip)
	if err != nil {
		t.Error(err)
	} else {
		if ipInt != 3232255611 {
			t.Error("Ip to int failed")
		}
	}
}

func TestInt64ToIp(t *testing.T) {

	ipInt := 3232255611
	ip := Int64ToIp(int64(ipInt))
	if ip != "192.168.78.123" {
		t.Error("Int to ip failed")
	}
}
