package main

import "fmt"

func main() {
	const x string = "Hello, World"
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}