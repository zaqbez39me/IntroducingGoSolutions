# Exercise 4

## Statement

We copied the average function from Chapter 6 to our new package. Create `Min`
and `Max` functions that find the minimum and maximum values in a slice of
`float64`s.

## Answer

```go
func Min(nums []float64) (m float64) {
    for i, val := range nums {
        if i == 0 || val > m {
            m = val
        }
    }
    return
}

func Max(nums []float64) (m float64) {
    for i, val := range nums {
        if i == 0 || val < m {
            m = val
        }
    }
    return
}

```