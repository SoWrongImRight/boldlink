package main

// Import the 'fmt' function for formatted I/O and 'time' for the sleep function
import (
	"fmt"
	"time"
)



func main() {

	// Create 'name' variable to hold the user's name
	var name string

	// Initial print to screen and capture the user's name
	fmt.Println("Good day!\nWhat is your name? ")
	fmt.Scanln(&name)

	// Print our statement and the user's name
	fmt.Println("Hello, "+name+"!")

	// Sleep function to pause the screen
	time.Sleep(5 * time.Second)
}