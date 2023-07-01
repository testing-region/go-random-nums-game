package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// check errors
func checkIfInt(err error) {
	if err != nil {
		log.Panicln("Please enter an integer")
	}
}

// get option
func getOption() (int, error) {
	option := bufio.NewScanner(os.Stdin)
	option.Scan()
	choice, err := strconv.ParseInt(option.Text(), 10, 64)
	return int(choice), err
}

// game levels
func pickLevel() (int, string) {
	var guess int
	var prompt string

	// display options menu
	fmt.Println()
	fmt.Println("Choose a level")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("1. Easy [0 - 10]")
	fmt.Println("2. Medium [0 - 25]")
	fmt.Println("3. Hard [0 - 50]")
	fmt.Println("4. Expert [0 - 100]")
	fmt.Println("5. Legendary [0 - 1000]")
	fmt.Println(strings.Repeat("-", 25))

	// get input
	fmt.Print("\noption: ")
	level, err := getOption()

	checkIfInt(err)

	// get the random number to guess and its corresponding prompt
	switch level {
	case 1:
		guess = rand.Intn(10)
		prompt = "Guess a number between 0 and 10: "
	case 2:
		guess = rand.Intn(25)
		prompt = "Guess a number between 0 and 25: "
	case 3:
		guess = rand.Intn(50)
		prompt = "Guess a number between 0 and 50: "
	case 4:
		guess = rand.Intn(100)
		prompt = "Guess a number between 0 and 100: "
	case 5:
		guess = rand.Intn(1000)
		prompt = "Guess a number between 0 and 1000: "
	}

	return guess, prompt
}

// returns true if player wins a game
func start(guess int, prompt string) bool {

	for tries := 1; tries < 6; tries++ {
		fmt.Print(prompt)
		userGuess, err := getOption()
		checkIfInt(err)
		fmt.Println()

		switch {

		case tries == 5:
			if userGuess == guess {
				fmt.Println("You guessed correctly, Good job!")
				return true
			}

		default:
			if userGuess < guess {
				fmt.Println("Too low, try again!")
			} else if userGuess > guess {
				fmt.Println("Too high, try again!")
			} else {
				fmt.Println("You guessed correctly, Good job!")
				if tries == 1 {
					fmt.Println()
					fmt.Println("Wow! You are amazing.")
					fmt.Println("You guessed the number correctly on your first attempt!")
				} else {
					fmt.Println("It took you", tries, "tries")
				}
				return true
			}

		}
	}

	fmt.Println("You lost, The number is", guess)
	fmt.Println("Better luck next time.")
	return false
}

func main() {
	var gamesWon, gamesPlayed int

	// Display message on startup
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Guess The Number")
	fmt.Println("Let's see how good you are at this")
	fmt.Println("[!] You have 5 tries, Good Luck!!!")
	fmt.Println(strings.Repeat("-", 40))

	for true {
		guess, prompt := pickLevel()
		win := start(guess, prompt)
		gamesPlayed++

		if win {
			gamesWon++
		}

		fmt.Println()
		fmt.Print("Do you want to play again? [1/0]: ")
		option, err := getOption()
		checkIfInt(err)

		if option != 1 {
			break
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Results")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("You played", gamesPlayed, "game(s).")
	fmt.Println("You won", gamesWon, "game(s), hope you had fun.")

	fmt.Println()
	fmt.Println("Sorry to see you go...")
	fmt.Println("Thanks for playing!!!")
}
