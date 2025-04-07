package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generating a random number between 1 and 100
	target := random.Intn(100)

	//Welcome message
	fmt.Println("Welcome to Guessing Game!")
	fmt.Println("I have choosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		//check guess equal to target
		if guess == target {
			fmt.Println("You win!")
			break
		}

		//check guess smaller than target
		if guess < target {
			fmt.Println("My number is bigger than", guess)
			continue
		}

		//check guess larger than target
		if guess > target {
			fmt.Println("My number is smaller than", guess)
			continue
		}
	}
}
