package medium_0347

// 建立一个堆
// 构建小顶堆
func buildHeap(nums []int, occurence map[int]int) {
	// 开始建立堆
	for i := len(nums) / 2; i >= 0; i-- {
		// 判断左子树
		if i*2+1 < len(nums) && occurence[nums[i]] > occurence[nums[2*i+1]] {
			tmp := nums[i]
			nums[i] = nums[2*i+1]
			nums[2*i+1] = tmp
		}
		// 判断右子树
		if i*2+2 < len(nums) && occurence[nums[i]] > occurence[nums[2*i+2]] {
			tmp := nums[i]
			nums[i] = nums[2*i+2]
			nums[2*i+2] = tmp
		}
	}
}

// 下滤操作
func filterDown(nums []int, occurence map[int]int) {
	index := 0
	for index < len(nums) {
		// 左子树右子树都有
		if index*2+2 < len(nums) {
			target := index
			// 找到左右子树中较小的数字与当前进行交换
			// 这样才能尽量少或者不破坏当前的堆栈的结构
			cindex := index*2 + 1
			if occurence[nums[cindex]] > occurence[nums[index*2+2]] {
				cindex = index*2 + 2
			}
			if occurence[nums[index]] > occurence[nums[cindex]] {
				tmp := nums[index]
				nums[index] = nums[cindex]
				nums[cindex] = tmp
				target = cindex
			}
			// 如果此时已经在正确的位置上,则不再进行替换操作
			if target == index {
				break
			} else {
				index = target
			}
		} else if index*2+1 < len(nums) {
			// 此时只有左子树
			target := index
			// 比当前的数据大的话,则换上来
			if occurence[nums[index]] > occurence[nums[index*2+1]] {
				tmp := nums[index]
				nums[index] = nums[index*2+1]
				nums[index*2+1] = tmp
				target = index*2 + 1
			}
			// 如果此时已经在正确的位置上,则不再进行替换操作
			if target == index {
				break
			} else {
				index = target
			}
		} else {
			// 没有孩子节点,则停止操作
			break
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) > 0 && k > 0 {
		// 先统计出一共又多少数据
		occurences := map[int]int{}
		for _, val := range nums {
			if _, ok := occurences[val]; ok {
				occurences[val]++
			} else if !ok {
				occurences[val] = 1
			}
		}
		// 建heap
		result := []int{}
		for key, val := range occurences {
			if len(result) < k {
				result = append(result, key)
				// 建堆操作
				buildHeap(result, occurences)
			} else {
				// 对堆进行调整
				if occurences[result[0]] < val {
					result[0] = key
					filterDown(result, occurences)
				}
			}
		}
		return result
	}
	return []int{}
}
