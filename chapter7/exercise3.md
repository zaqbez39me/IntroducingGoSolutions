# Exercise 3

## Statement

Add a new perimeter method to the Shape interface to calculate the perimeter of
a shape. Implement the method for Circle and Rectangle.

## Answer

```go
func (r *Rectangle) perimeter() float64 {
	return 2*math.Abs(r.x2-r.x1) + 2*math.Abs(r.y2-r.y1)
}

func (c *Circle) perimeter() float64 {
	return 2 * c.r * math.Pi
}
```