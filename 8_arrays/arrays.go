package main

import (
	"fmt"
)

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	fmt.Println(len(nums))

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums2)

	nums2d := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums2d)
}
