package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
	usage    = `
Usage: %d <guess>
Your misson is to guess one of those numbers
The greater your guess, the harder it is
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please enter a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose!")
}
