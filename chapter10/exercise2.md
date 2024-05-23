# Exercise 2

## Statement

Write your own `Sleep` function using `time.After`.

## Answer

```go
func Sleep(duration time.Duration) {
	for {
		select {
		case <-time.After(duration):
			return
		}
	}
}
```
