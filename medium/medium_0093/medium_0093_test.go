package medium_0093

import (
	"reflect"
	"testing"
)

// success
func TestRestoreIpAddresses(t *testing.T) {

	var restoreIpAddressesTest = []struct {
		in       string   // input
		expected []string // expected result
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"010010", []string{"0.10.0.10", "0.100.1.0"}},
	}

	for _, riat := range restoreIpAddressesTest {
		actual := restoreIpAddresses(riat.in)
		if !reflect.DeepEqual(actual, riat.expected) {
			t.Errorf("restoreIpAddressesTest(%s) = %s; expected %s", riat.in, actual, riat.expected)
		}
	}
}
