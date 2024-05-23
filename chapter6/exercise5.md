# Exercise 5

## Statement

The Fibonacci sequence is defined as: `fib(0) = 0`, `fib(1) = 1`, `fib(n) =
fib(n-1) + fib(n-2)`. Write a recursive function that can find `fib(n)`.

## Solution

```go
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
```