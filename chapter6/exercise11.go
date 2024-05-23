package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
	return
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("x=%d and y=%d", x, y)
}
