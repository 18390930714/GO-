package main

import "fmt"

//冒泡排序:

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(bubblesort(arr))
}

func bubblesort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			fmt.Println(nums)
		}
		fmt.Println("**************************")
	}
	return nums
}
