package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// channel
// can only pass one type to channel
var keyPressChan chan rune

func main() {
	// 3-5. channels
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit.")

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		// send char to channel
		keyPressChan <- char
	}
}

// channels

func listenForKeyPress() {
	for {
		key := <-keyPressChan // key is whatever passed into keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
