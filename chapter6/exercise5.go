package main

import "fmt"

var dp = make(map[uint]uint)

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if val, ok := dp[n]; ok {
		return val
	}
	dp[n] = fib(n-1) + fib(n-2)
	return dp[n]
}

func main() {
	fmt.Printf("Input the number of Fibonacci sequence: ")
	var n uint
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	fmt.Println(fib(n))
}
