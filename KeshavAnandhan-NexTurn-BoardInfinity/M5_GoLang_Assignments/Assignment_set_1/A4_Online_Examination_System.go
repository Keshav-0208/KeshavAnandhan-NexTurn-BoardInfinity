package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type question struct {
	questionText string
	options      []string
	answerIndex  int
}

var questions = []question{
	{
		questionText: "What does OOP stand for?",
		options:      []string{"1. Object-Oriented", "2. Open Operator", "3. Online Option", "4. Object Operator"},
		answerIndex:  1,
	},
	{
		questionText: "Which language is used for web?",
		options:      []string{"1. Go", "2. C", "3. Java", "4. JavaScript"},
		answerIndex:  4,
	},
	{
		questionText: "What is Git for?",
		options:      []string{"1. Code management", "2. Web design", "3. Database", "4. Networking"},
		answerIndex:  1,
	},
}

func startQuiz() (int, error) {
	score := 0
	fmt.Println("Quiz Start!")
	fmt.Println("Type 'quit' to exit.\n")

	for i := 0; i < len(questions); i++ {
		fmt.Println("Q", i+1, ":", questions[i].questionText)
		for _, option := range questions[i].options {
			fmt.Println(option)
		}
		fmt.Print("Your answer: ")

		inputChannel := make(chan string)
		go func() {
			var input string
			fmt.Scan(&input)
			inputChannel <- input
		}()

		select {
		case userInput := <-inputChannel:
			if strings.ToLower(userInput) == "quit" {
				return score, nil
			}

			choice, err := validateInput(userInput, len(questions[i].options))
			if err != nil {
				fmt.Println("Invalid input, try again.")
				i-- // Retry question
				continue
			}

			if choice == questions[i].answerIndex {
				score++
			}
		case <-time.After(10 * time.Second): // Time out after 10 seconds
			fmt.Println("\nTime's up!")
		}
	}
	return score, nil
}

func validateInput(input string, numOptions int) (int, error) {
	var choice int
	_, err := fmt.Sscan(input, &choice)
	if err != nil || choice < 1 || choice > numOptions {
		return 0, errors.New("invalid choice")
	}
	return choice, nil
}

func main() {
	totalQuestions := len(questions)

	finalScore, err := startQuiz()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nQuiz Finished!")
	fmt.Println("Score:", finalScore, "/", totalQuestions)
}
