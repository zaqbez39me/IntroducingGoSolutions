package main

import "fmt"

// this is a comment
func main() {
	var name string
	fmt.Printf("Input your name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return
	}
	fmt.Println("Hello, my name is", name)
}
