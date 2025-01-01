package main

import "fmt"

func toFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func toCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var temp float64
	var unit string

	fmt.Print("Enter temperature: ")
	fmt.Scan(&temp)
	fmt.Print("Is this in Celsius or Fahrenheit? : ")
	fmt.Scan(&unit)

	if unit == "C" || unit == "c" || unit == "Celsius" || unit == "celsius" {
		fmt.Printf("%.2f°C is %.2f°F\n", temp, toFahrenheit(temp))
	} else if unit == "F" || unit == "f" || unit == "Fahrenheit" || unit == "fahrenheit" {
		fmt.Printf("%.2f°F is %.2f°C\n", temp, toCelsius(temp))
	} else {
		fmt.Println("Invalid unit.")
	}
}
