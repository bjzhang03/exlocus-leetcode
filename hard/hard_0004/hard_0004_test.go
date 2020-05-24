package hard_0004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {

	var fmsa = []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, val := range fmsa {
		actual := findMedianSortedArrays(val.nums1, val.nums2)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %f, expected : %f", actual, val.expected)
		}
	}

}
