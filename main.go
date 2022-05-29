package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

// getNumber gets a number from a user
func getNumber() (int, error) {
	var num int
	// get an int value from a user
	_, err := fmt.Scanln(&num)

	if err != nil {
		// clear input stream when an error occurs
		stdin.ReadString('\n')
	}

	return num, err
}

func main() {
	// set a random seed(int64) to time
	rand.Seed(time.Now().UnixNano())

	// get a ramdon number(int), range 100 (0 to 99)
	rNum := rand.Intn(100)
	cnt := 1

	for {
		fmt.Print("Guess the number (0-99): ")
		num, err := getNumber()

		if err != nil {
			fmt.Println("Please guess the interger number.")
		} else {
			if num > 99 || num < 0 {
				fmt.Println("Please guess the number between 0 and 99, inclusive.")
			} else if rNum > num {
				fmt.Println("The number is greater than your guess.")
			} else if rNum < num {
				fmt.Println("The number is less than your guess.")
			} else {
				fmt.Println("Congrats! You guessed the right number on", cnt, "attempt(s).")
				break
			}
			cnt++
		}
	}
}
