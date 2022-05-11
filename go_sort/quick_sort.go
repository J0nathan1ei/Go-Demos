package main

import (
	"fmt"
)

func qSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	low, high := 0, length-1
	pivot := low

	for low < high {
		for low < high && nums[high] > nums[pivot] {
			high--
		}

		for low < high && nums[low] < nums[pivot] {
			low++
		}

		nums[low], nums[high] = nums[high], nums[low]
	}
	qSort(nums[:low])
	qSort(nums[low+1:])
}

func main() {
	d := []int{5, 9, 3, 7, 1, 8, 6, 4, 2, 0}
	qSort(d)
	fmt.Println(d)
}
