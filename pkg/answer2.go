package pkg

import (
	"math/rand"
)

//2. 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。

// 此题疑惑
// 扑克的位置可以被2个以上覆盖？
// 假设输入的扑克是一个混序的数组；那么需要输出一个新的均匀分布的数组？
// 还是将输入的数组，将其随机排列为不重叠的新数组？

// 选择将输入的数组随机在一个新的数组中进行随机排列，位置不重复；
func test(arr []string) []string {
	result := make([]string, len(arr))

	nums := gennerateRandom(0, len(arr), len(arr))

	for i, v := range nums {
		result[v] = arr[i]
	}
	return result
}

// 生成位置随机数组
func gennerateRandom(start, end, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	//生成随机数

	for len(nums) < count {
		//生成随机数
		num := rand.Intn(end)

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums

}