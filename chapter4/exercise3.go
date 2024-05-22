package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := make([]string, 0, 100)
	for i := 1; i <= 100; i++ {
		var currentValue string
		if i%15 == 0 { // Divisible by 3 and 5
			currentValue = "FizzBuzz."
		} else if i%5 == 0 { // Divisible only by 5
			currentValue = "Buzz."
		} else if i%3 == 0 { // Divisible only by 3
			currentValue = "Fizz."
		} else { // Not divisible by any of given numbers
			currentValue = fmt.Sprintf("%d", i)
		}
		numbers = append(numbers, currentValue)
	}
	fmt.Println(strings.Join(numbers, ","))
}
