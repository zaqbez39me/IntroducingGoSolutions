package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	halvedValue, isEven := half(1)
	fmt.Println("half(1) returns", fmt.Sprintf("(%d, %t)", halvedValue, isEven))
	halvedValue, isEven = half(2)
	fmt.Println("half(2) returns", fmt.Sprintf("(%d, %t)", halvedValue, isEven))
}
