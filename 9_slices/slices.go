package main

func main() {
	// Example usage of slices in Go

	// Uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// Create a slice with make
	// var nums = make([]int, 0, 5)
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// Empty slice
	// nums := []int{}

	// Append elements
	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// Modify elements
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// Copy slices
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// Slice operator
	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// Compare slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}
	// fmt.Println(slices.Equal(nums1, nums2))

	// Multi-dimensional slice
	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)
}
