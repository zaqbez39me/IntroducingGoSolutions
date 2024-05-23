# Exercise 3

## Statement

Write a function with one variadic parameter that finds the greatest number in a list of numbers.

## Solution

```go
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
```