package easy_0350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > 0 && len(nums2) > 0 {
		// 建立map来进行计数操作
		savenums1 := make(map[int]int)
		savenums2 := make(map[int]int)
		for i := 0; i < len(nums1); i++ {
			if _, ok := savenums1[nums1[i]]; ok {
				savenums1[nums1[i]]++
			} else {
				savenums1[nums1[i]] = 1
			}
		}
		for i := 0; i < len(nums2); i++ {
			if _, ok := savenums2[nums2[i]]; ok {
				savenums2[nums2[i]]++
			} else {
				savenums2[nums2[i]] = 1
			}
		}
		result := []int{}
		for k, v := range savenums1 {
			if _, ok := savenums2[k]; ok {
				count := v
				if count > savenums2[k] {
					count = savenums2[k]
				}

				for i := 0; i < count; i++ {
					result = append(result, k)
				}
			}
		}
		return result
	}
	return []int{}
}
