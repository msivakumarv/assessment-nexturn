package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct to represent a quiz question
type Question struct {
	Text       string
	Options    [4]string
	CorrectAns int
}

var questionBank = []Question{
	{
		Text:       "What is the capital of France?",
		Options:    [4]string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
		CorrectAns: 3,
	},
	{
		Text:       "What is 2 + 2?",
		Options:    [4]string{"1. 3", "2. 4", "3. 5", "4. 6"},
		CorrectAns: 2,
	},
	{
		Text:       "Which programming language is this quiz written in?",
		Options:    [4]string{"1. Python", "2. Go", "3. Java", "4. C++"},
		CorrectAns: 2,
	},
}

// Constants for quiz performance classification
const (
	Excellent        = "Excellent"
	Good             = "Good"
	NeedsImprovement = "Needs Improvement"
	TimeLimit        = 10 // seconds per question
)

// Function to start the quiz
func takeQuiz() (int, error) {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Text)
		for _, option := range question.Options {
			fmt.Println(option)
		}
		fmt.Printf("Enter your answer (1-4) or type 'exit' to quit: ")

		// Setting a timer for the question
		answerChan := make(chan string)
		go func() {
			answer, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(answer)
		}()

		select {
		case answer := <-answerChan:
			if strings.ToLower(answer) == "exit" {
				fmt.Println("Exiting the quiz early.")
				return score, nil
			}

			// Validate input
			selectedOption, err := strconv.Atoi(answer)
			if err != nil || selectedOption < 1 || selectedOption > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			// Check correctness
			if selectedOption == question.CorrectAns {
				score++
			}

		case <-time.After(TimeLimit * time.Second):
			fmt.Println("\nTime's up for this question! Moving to the next.")
			continue
		}
	}

	return score, nil
}

// Function to classify performance
func classifyPerformance(score int, totalQuestions int) string {
	percentage := (float64(score) / float64(totalQuestions)) * 100
	switch {
	case percentage >= 80:
		return Excellent
	case percentage >= 50:
		return Good
	default:
		return NeedsImprovement
	}
}

func main() {
	fmt.Println("Welcome to the Online Quiz System!")
	fmt.Println("You have", TimeLimit, "seconds to answer each question.")

	totalQuestions := len(questionBank)
	score, err := takeQuiz()
	if err != nil {
		fmt.Println("Error during the quiz:", err)
		return
	}

	fmt.Printf("\nYou completed the quiz! Your score: %d/%d\n", score, totalQuestions)
	performance := classifyPerformance(score, totalQuestions)
	fmt.Println("Performance:", performance)
}
