package pkg

//实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出，下一个人接着从1开始报数，如此反复，求最后的胜利者编号。
//例如，n=3，m=2，那么失败者编号依次是1、0，最后的胜利者是2号。
//这里考虑m，n都是正常的数据范围，其中：
//1 <= n <= 10^5
//1 <= m <= 10^6
//算法要求考虑时间效率。



func Game(n, m int) int {
	if n == 1 {
		return n
	}

	//m := map[string]int[]{}
	respArr := make([]int, 0)
	for i:=0;i<n;i++{
		respArr = append(respArr, i)
	}

	// 时间复杂度 O(2^n)
	resp := calculate(respArr, m, 1)
	return resp[0]


	// 时间复杂度O(n^2)
	//num := 1
	//for {
	//	if len(respArr) == 1 {
	//		return respArr[0]
	//	}
	//	resp1 := make([]int, 0)
	//	for _, v := range respArr{
	//		if num == m {
	//			num = 1
	//		} else {
	//			respArr = append(resp1, v)
	//			num++
	//		}
	//	}
	//}
}


// 时间复杂度 O（2*n）
func calculate(arr []int, m, num int) []int {
	if len(arr) == 1 {
		return arr
	}
	respArr := make([]int, 0)
	for _, v := range arr{
		if num == m {
			num = 1
		} else {
			respArr = append(respArr, v)
			num++
		}
	}
	return calculate(respArr, m, num)
}