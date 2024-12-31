package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number to calculate its factorial: ")
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Factorial is not defined for negative numbers.")
        return
    }

    factorial := 1
    for i := 1; i <= num; i++ {
        factorial *= i
    }

    fmt.Printf("The factorial of %d is %d.\n", num, factorial)
}
