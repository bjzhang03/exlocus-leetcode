package medium_0468

import "testing"

func TestValidIPAddress(t *testing.T) {

	var vips = []struct {
		IP       string
		expected string
	}{
		{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"1e1.4.5.6", "Neither"},
		{"20EE:Fb8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"192.0.0.1", "IPv4"},
	}

	for _, val := range vips {
		actual := validIPAddress(val.IP)
		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}
}
