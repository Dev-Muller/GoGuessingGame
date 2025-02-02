package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A number will be generated between 1 and 100, try to guess it!")

	x := int64(rand.Intn(101))
	x = 10
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Enter your guess:")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessesInt, err := strconv.ParseInt(guess, 10, 64)

		if err != nil {
			fmt.Println("Please enter a valid number")
			return
		}

		switch {
		case guessesInt < x:
			fmt.Println("You are wrong. Too low", guessesInt)
		case guessesInt > x:
			fmt.Println("You are wrong. Too high", guessesInt)
		case guessesInt == x:
			fmt.Printf("You are correct! The number was: %d\n"+
				"You guessed right in %d guesses\n"+
				"These were your guesses: %v",
				x, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessesInt
	}

	fmt.Printf(
		"You have reached the maximum number of guesses. The number was: %d\n"+
			"You had 10 guesses\n"+
			"These were your guesses: %v",
		x, guesses,
	)
}
