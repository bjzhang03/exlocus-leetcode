package medium_0406

// https://blog.csdn.net/mengyujia1234/article/details/89646569
// 这个题目我看了很久都没有看出来解法,然后看了别人的解法
func reconstructQueue(people [][]int) [][]int {
	if len(people) > 0 {
		result := [][]int{}
		for len(people) > 0 {
			// 找到当前的最大的数据
			currentm := getMax(people)
			// 将数据分成两部分,一部分是当前数据里面值最大的数字,另一部分不是最大值
			np := [][]int{}
			sp := [][]int{}
			for i := 0; i < len(people); i++ {
				if people[i][0] == currentm {
					sp = append(sp, people[i])
				} else {
					np = append(np, people[i])
				}
			}

			// 将当前的数据最大的值插入到结果中
			insertItem(&result, sp)
			people = people[0:0]
			people = append(people, np...)
		}
		return result
	}
	return [][]int{}

}

func insertItem(save *[][]int, items [][]int) {
	for len(items) > 0 {
		// 找到item中最小的数据
		sti := 0
		tmp := items[0]
		for i := 0; i < len(items); i++ {
			if tmp[1] > items[i][1] {
				sti = i
				tmp = append([]int{}, items[i]...)
			}
		}
		// 将数据插入进来
		right := append([][]int{}, (*save)[tmp[1]:]...)
		*save = append((*save)[:tmp[1]], tmp)
		*save = append((*save), right...)
		// 一个数据
		items = append(items[:sti], items[sti+1:]...)
	}
}

// 获取数据的最大值
func getMax(people [][]int) int {
	result := 0
	for i := 0; i < len(people); i++ {
		if result < people[i][0] {
			result = people[i][0]
		}
	}
	return result
}
