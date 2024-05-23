# Exercise 4

## Statement

Using makeEvenGenerator as an example, write a makeOddGenerator function
that generates odd numbers.

## Solution

```go
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
```