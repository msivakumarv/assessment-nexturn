package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getHint(target, guess int) string {
	if guess > target {
		return "Too high!"
	} else if guess < target {
		return "Too low!"
	}
	return "Correct!"
}

func main() {
	var userGuess int
	var maxAttempts = 5
	var attempts int

	rand.Seed(time.Now().UnixNano())
	targetNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Guess a number between 1 and 100.")
	fmt.Printf("You have %d attempts to guess correctly.\n", maxAttempts)

	for attempts < maxAttempts {
		fmt.Printf("\nAttempt %d: Enter your guess: ", attempts+1)
		_, err := fmt.Scan(&userGuess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		hint := getHint(targetNumber, userGuess)
		fmt.Println(hint)

		if hint == "Correct!" {
			fmt.Println("Congratulations! You guessed the correct number!")
			return
		}
		attempts++
	}

	fmt.Printf("\nYou've used all your attempts! The correct number was %d.\n", targetNumber)
	fmt.Println("Better luck next time!")
}
