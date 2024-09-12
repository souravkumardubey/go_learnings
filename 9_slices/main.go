package main

import (
	"fmt"
	"slices"
)

// ? slice -> dynamic array in Go
func main() {

	// * uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)

	// * 2nd parameter defines the initial length
	// nums := make([]int, 2);
	// fmt.Println(nums)

	// * 3rd parameter defines the increase in size once it is full
	// nums := make([]int, 0, 5) // * initial capacity is 5, increases by 5 if full
	// * another way to do it
	// nums := []int {}
	// nums = append(nums, 1);
	// nums = append(nums, 2);
	// nums = append(nums, 3);
	// nums[2] = 2
	// fmt.Println(nums[2])
	// fmt.Println(cap(nums))


	// * copy()
	// nums := make([]int, 0, 5)
	// nums = append(nums, 1)
	// nums2 := make([]int, len(nums))
	// copy(nums2, nums);

	// fmt.Println(nums, nums2)
	// fmt.Println(len(nums2))
	// fmt.Println(cap(nums2))

	// * slice operator
	// var nums = []int{0, 1, 2}
	// fmt.Println((nums[0:2])) // or nums[:2]
	
	// * slice
	var nums1 = []int{0, 1, 2}
	var nums2 = []int{0, 1, 2}
	
	fmt.Println(slices.Equal(nums1, nums2));
	

}