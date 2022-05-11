package main

import (
	"fmt"
)

func adjustHeap(nums []int, parent, end int) {
	child := parent*2 + 1
	for child < end {
		if child+1 < end && nums[child+1] > nums[child] {
			child++
		}

		if nums[parent] > nums[child] {
			return
		}

		nums[parent], nums[child] = nums[child], nums[parent]
		parent = child
		child = child*2 + 1
	}
}

func buildHeap(nums []int) {
	length := len(nums)
	parent := length/2 - 1
	for i := parent; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}
}

func heapSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	buildHeap(nums)
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i-1)
	}

}

func main() {
	d := []int{5, 9, 3, 7, 1, 8, 6, 4, 2, 0}
	heapSort(d)
	fmt.Println(d)
}
