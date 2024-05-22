package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := make([]string, 0, 33)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			numbers = append(numbers, fmt.Sprintf("%d", i))
		}
	}
	fmt.Println(strings.Join(numbers, ","))
}
