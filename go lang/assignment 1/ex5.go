package main

import "fmt"

func main() {
	var name string
	var age int
	var favoriteNumber float64

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter your favorite number: ")
	fmt.Scanln(&favoriteNumber)

	fmt.Printf("Hello, %s! You are %d years old, and your favorite number is %.2f.\n", name, age, favoriteNumber)
}
