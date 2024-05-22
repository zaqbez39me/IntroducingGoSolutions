package main

import "fmt"

func main() {
	var inputFeet float64
	fmt.Printf("Input the feets to convert: ")
	_, err := fmt.Scanf("%f", &inputFeet)
	if err != nil {
		return
	}
	fmt.Println("Converted value in meters:", inputFeet*0.3048) // 1 ft = 0.3048 m
}
