package main

import "fmt"

// 选择排序: 选一个数当最小值或者最大值 然后进行比较交换，循环向后查询
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(selectsort(arr))
}

func selectsort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
