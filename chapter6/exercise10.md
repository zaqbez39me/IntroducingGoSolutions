# Exercise 10

## Statement

What is the value of x after running this program:

```go
func square(x *float64) {
*x = *x * *x
}
func main() {
x := 1.5
square(&x)
}
```

## Solution

In the beginning of the program the value of x is equal to `1.5`. Then the value of x is updated inside the square
function by square of 1.5 what equals to `2.25`.

## Answer

2.25

