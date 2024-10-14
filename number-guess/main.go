package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rand_number := rand.IntN(100)
	guess := 0

	for {
		fmt.Println("Guess a number between 0 and 100: ")
		fmt.Scanf("%d", &guess)
		if guess < 0 || guess > 100 {
			fmt.Println("Enter an integer between 0 and 100")
		}
		if guess < rand_number {
			fmt.Println("Too low")
		} else if guess > rand_number {
			fmt.Println("Too high")
		} else if guess == rand_number {
			fmt.Println("Correct!")
			fmt.Println("The number was: ", rand_number)
			break
		}
	}

}
