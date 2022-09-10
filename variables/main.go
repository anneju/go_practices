package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// one way, declare, then assign (two steps)
	var firstNumber int
	firstNumber = rand.Intn(8) + 2
	// another way, declare type and name and assign value
	var secondNumber = rand.Intn(8) + 2
	// one step variable: declare name, assign value, and let Go fighure out the type
	subtraction := rand.Intn(8) + 2

	fmt.Println("First number", firstNumber)
	fmt.Println("2nd number", secondNumber)
	fmt.Println("3rd number", subtraction)

	// store the answer
	var answer int = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// create our reader variable, which reads input from standard in
	reader := bufio.NewReader(os.Stdin)

	// dispaly a welcome/instructions
	fmt.Println("Guess the number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 - 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply your number by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you origionally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	// give them the answer
	// answer = firstNumber*secondNumber - subtraction
	fmt.Println("The asnwer is", answer)
}
