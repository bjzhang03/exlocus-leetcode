package easy_0069

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	var ms = []struct {
		x        int
		expected int
	}{
		{4, 2},
		{8, 2},
	}

	for _, val := range ms {
		actual := mySqrt(val.x)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
