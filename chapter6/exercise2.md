# Exercise 2

## Statement

Write a function that takes an integer and halves it and returns true if it was even
or false if it was odd. For example, half(1) should return (0, false) and
half(2) should return (1, true).

## Solution

```go
func half(n int) (int, bool) {
return n / 2, n%2 == 0
}
```
