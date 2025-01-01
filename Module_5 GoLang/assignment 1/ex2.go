package main

import "fmt"

func main() {
	var side1, side2, side3 int
	fmt.Print("Enter side 1: ")
	fmt.Scan(&side1)
	fmt.Print("Enter side 2: ")
	fmt.Scan(&side2)
	fmt.Print("Enter side 3: ")
	fmt.Scan(&side3)

	if side1 <= 0 || side2 <= 0 || side3 <= 0 {
		fmt.Println("Sides must be positive integers.")
		return
	}

	if side1+side2 <= side3 || side1+side3 <= side2 || side2+side3 <= side1 {
		fmt.Println("The entered sides do not form a valid triangle.")
		return
	}

	if side1 == side2 && side2 == side3 {
		fmt.Println("The triangle is Equilateral.")
	} else if side1 == side2 || side2 == side3 || side1 == side3 {
		fmt.Println("The triangle is Isosceles.")
	} else {
		fmt.Println("The triangle is Scalene.")
	}
}