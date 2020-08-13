package hard_0004

import "fmt"

// https://www.jianshu.com/p/788555ce7a51
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > 0 && len(nums2) > 0 {
		if target := (len(nums1) + len(nums2)) / 2; (len(nums1)+len(nums2))%2 == 0 {
			fmt.Println("target :=", target)
			sta1, sta2, end1, end2 := 0, 0, len(nums1)-1, len(nums2)-1
			for sta1 <= end1 && sta2 <= end2 {
				fmt.Println(nums1, sta1, nums2, sta2)
				if sta1+sta2 == target {
					return float64(nums1[sta1]+nums2[sta2]) / 2
				}
				mid1 := (sta1 + end1) / 2
				mid2 := (sta2 + end2) / 2
				if nums1[mid1] < nums2[mid2] {
					sta1 = mid1 + 1
				} else if nums1[mid1] >= nums2[mid2] {
					sta2 = mid2 + 1
				}
			}
		} else {
			sta1, sta2, end1, end2 := 0, 0, len(nums1)-1, len(nums2)-1
			for sta1 <= end1 && sta2 <= end2 {
				if sta1+sta2 == target {
					if nums1[sta1] < nums2[sta2] {
						return float64(nums1[sta1])
					} else {
						return float64(nums2[sta2])
					}
				}
				mid1 := (sta1 + end1) / 2
				mid2 := (sta2 + end2) / 2
				if nums1[mid1] < nums2[mid2] {
					sta1 = mid1 + 1
				} else if nums1[mid1] >= nums2[mid2] {
					sta2 = mid2 + 1
				}

			}

		}
	} else if len(nums1) > 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2]) / 2
		} else {
			return float64(nums1[len(nums1)/2])
		}
	} else if len(nums2) > 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2]) / 2
		} else {
			return float64(nums2[len(nums2)/2])
		}
	}
	return 0
}
