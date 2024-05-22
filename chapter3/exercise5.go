package main

import "fmt"

func main() {
	var fahrenheitInput float64
	fmt.Printf("Input the value in Farnheit: ")
	// Read the value in Fahrenheit and capture err
	_, err := fmt.Scanf("%f", &fahrenheitInput)
	if err != nil {
		return
	}
	fmt.Printf("Corresponding value in Celsius: %f", (fahrenheitInput-32)/9*5) // C = (F − 32) * 5/9
}
