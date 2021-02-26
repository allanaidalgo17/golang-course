package main

import (
	"fmt"
	"time"
)

func speak(name, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (i = %d)\n", name, text, i)
	}
}

func main() {
	// speak("Person1", "Text 1", 3)
	// speak("Person2", "Text 2", 1)

	// Goroutine
	// a goroutine only runs while the main program is running
	go speak("Goroutine", "Hi!", 10)
	speak("Serial", "Hello!", 5)
}
