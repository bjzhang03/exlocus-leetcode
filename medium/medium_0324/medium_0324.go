package medium_0324

import (
	"fmt"
)

// 建立一个堆
func buildHeap(nums []int) {
	// 开始建立堆
	for i := len(nums) / 2; i >= 0; i-- {
		if i*2 < len(nums) && nums[i] <= nums[2*i] {
			tmp := nums[i]
			nums[i] = nums[2*i]
			nums[2*i] = tmp
		}

		if i*2+1 < len(nums) && nums[i] <= nums[2*i+1] {
			tmp := nums[i]
			nums[i] = nums[2*i+1]
			nums[2*i+1] = tmp
		}
	}
}

// 下滤操作
func filterDown(nums []int) {
	// 新建临时的存储
	heapSave := []int{0}
	heapSave = append(heapSave, nums...)
	index := 1

	for index < len(heapSave) {
		// 判断叶子节点
		if index*2+1 < len(heapSave) {
			target := index
			// 比当前的数据大的话,则换上来
			if heapSave[index] < heapSave[index*2+1] {
				tmp := heapSave[index]
				heapSave[index] = heapSave[index*2+1]
				heapSave[index*2+1] = tmp
				target = index*2 + 1
			}
			// 比当前的数据大的话,则换上来
			if heapSave[index] < heapSave[index*2] {
				tmp := heapSave[index]
				heapSave[index] = heapSave[index*2]
				heapSave[index*2] = tmp
				target = index * 2
			}
			// 如果此时已经在正确的位置上,则不再进行替换操作
			if target == index {
				break
			} else {
				index = target
			}

		} else if index*2 < len(heapSave) {
			target := index
			// 比当前的数据大的话,则换上来
			if heapSave[index] < heapSave[index*2] {
				tmp := heapSave[index]
				heapSave[index] = heapSave[index*2]
				heapSave[index*2] = tmp
				target = index * 2
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
	// 清空数据
	nums = nums[0:0]
	nums = append(nums, heapSave[1:]...)
}

// 堆排序
// 上滤用于添加元素
// 下滤用于修改元素
func splitNums(nums []int, k int) ([]int, []int) {
	leftk := make([]int, k)
	right := []int{}

	for i := 0; i < k; i++ {
		leftk[i] = nums[i]
	}

	buildHeap(leftk)
	for i := k; i < len(nums); i++ {
		if leftk[0] >= nums[i] {
			right = append(right, leftk[0])
			leftk[0] = nums[i]
			filterDown(leftk)
		} else {
			right = append(right, nums[i])
		}
	}

	return leftk, right
}

// https://www.cnblogs.com/grandyang/p/5139057.html
func wiggleSort(nums []int) {
	if len(nums) > 0 {
		// 先对数据进行排序操作
		count := len(nums) / 2
		if len(nums)%2 != 0 {
			count = len(nums)/2 + 1
		}
		left, right := splitNums(nums, count)

		fmt.Println("split = ", left, right)

	}
}
