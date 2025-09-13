package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func rockPaperScissors() error {
	const rounds = 3

	fmt.Println("Let's play. You have", rounds, "rounds")

	//for loop for game rounds
	for i := 0; i < rounds; i++ {
		//generate computer choice
		computerChoiceNum := rand.Intn(rounds) //random number between 0-2
		var computerChoice string

		switch computerChoiceNum {
		case 0:
			computerChoice = "rock"
		case 1:
			computerChoice = "paper"
		case 2:
			computerChoice = "scissors"
		}

		// read playeres choice
		var playerChoice string
		fmt.Print("Enter your choice - rock paper or scissors")
		fmt.Scanln(&playerChoice)

		if !slices.Contains([]string{"rock", "paper", "scissors"}, playerChoice) {
			return errors.New("Invalid input")
		}

		//game logic
		fmt.Printf("Computer chose: %s\n", computerChoice)
		switch {
		case playerChoice == computerChoice:
			fmt.Println("tie")
		case playerChoice == "rock" && computerChoice == "scissors", playerChoice == "paper" && computerChoice == "rock", playerChoice == "scissors" && computerChoice == "paper":
			fmt.Println("You win")
		default:
			fmt.Println("Computer wins")
		}
	}

	fmt.Println("Game over")
	return nil
}

func main() {
	err := rockPaperScissors()
	if err != nil {
		fmt.Println("errror:", err)
	}
}
