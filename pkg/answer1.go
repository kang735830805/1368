package pkg

//1. 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。

func QuickSort(arr []int64)  {
	if len(arr) == 0 {
		return
	}
	mid, temp := arr[0], 1
	head, tail := 0, len(arr)-1
	for head < tail {
		if arr[temp] > mid {
			arr[temp], arr[tail] = arr[tail], arr[temp]
			tail--
		} else {
			arr[temp], arr[head] = arr[head], arr[temp]
			head++
			temp++
		}
	}
	arr[head]=mid
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
}

func Run1(arr []int64) int64 {
	QuickSort(arr)
	return arr[len(arr)-1]-arr[0]
}