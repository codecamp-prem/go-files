// guess challenges players to guess a random numbers
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Get current date time as Integers
	rand.Seed(seconds)           // Seed the random number generator
	target := rand.Intn(100) + 1 // generate an integer between 1 to 100
	fmt.Println("I've choosen a random number 1 to 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) //create a buffio.Reader, which lets read keyboard input
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10-guesses, " guesses left")
		fmt.Println("Make a guess")
		input, err := reader.ReadString('\n') // Reads what user types up until when they press Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("OOps! your guess is low")
		} else if guess > target {
			fmt.Println("OOps! your guess is High")
		} else {
			success = true
			fmt.Println("Congrats! You guessed correctly.")
			break
		}
	}
	if !success {
		fmt.Println("Sorry! you didn't guess correctly.")
	}
}
