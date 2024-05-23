package main

import "fmt"

func greatestNumber(nums ...int) (greatest int) {
	if len(nums) == 0 {
		panic("Given empty sequence")
	}
	greatest = nums[0]
	for _, num := range nums {
		if greatest < num {
			greatest = num
		}
	}
	return
}

func main() {
	fmt.Println(greatestNumber(1, 2, 3))
}
