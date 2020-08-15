package easy_0070

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var cs = []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
	}

	for _, val := range cs {
		actual := climbStairs(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
