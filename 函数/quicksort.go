package main

import "fmt"

//快速排序

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(quicksort(arr))
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, splitdata)
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = quicksort(low), quicksort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}
