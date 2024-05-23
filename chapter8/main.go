package main

import (
	m "IntroductionGo/chapter8/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
