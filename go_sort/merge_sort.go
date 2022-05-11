package main

import (
	"fmt"
)

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	var res []int
	mid := length / 2
	res = merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))

	return res
}

func merge(left, right []int) []int {
	var res []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		for _, v := range left {
			res = append(res, v)
		}
	}

	if len(right) > 0 {
		for _, v := range right {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	d := []int{5, 9, 3, 7, 1, 8, 6, 4, 2, 0}
	t := mergeSort(d)
	fmt.Println(t)
}
