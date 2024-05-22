# Exercise 1

## Statement

What does the following program print?

```go

i := 10
if i > 10 {
fmt.Println("Big")
} else {
fmt.Println("Small")
}

```

## Solution

Let us analyze the given program. First of all, it declares and initializes the variable `i` with value `10`. then it
checks whether the value of variable `i` is larger than 10. It is not larger than 10. Hence, we execute instructions
from else clause. We see only single instruction `fmt.Println("Small")` that simply prints `Small` in console.

## Answer

`Small`