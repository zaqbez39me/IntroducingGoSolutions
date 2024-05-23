# Exercise 1

## Statement

Writing a good suite of tests is not always easy, but the process of writing tests
often reveals more about a problem than you may at first realize. For example,
with our Average function, what happens if you pass in an empty list
`([]float64{})`? How could the function be modified to return 0 in this case?

## Answer

When we pass `([]float64{})` we run into runtime error. We can add check on zero length slice right in the beginning og
the function. Here is how:
```go
func Average(xs []float64) float64 {
    if len(xs) == 0 {
        return 0
    }
	... // The rest of the function body
}

```