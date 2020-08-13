package medium_0470

//https://www.jianshu.com/p/ba2bf153d1ad?from=singlemessage
// 这道题目是一个很经典的数学题,这里参考了别人的方案,记录一下
// 产生随机数难度很低,难度高的是要求所有的概率都是相等的
func rand10() int {
	for {
		result := (rand7()-1)*7 + rand7() - 1

		if result < 40 {
			return result%10 + 1
		}
	}
	return 0
}
