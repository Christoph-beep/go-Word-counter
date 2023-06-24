package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// inform user to enter input and get input from terminal

// user input
func main() {

	fmt.Println("Enter your words here")

	// reader instance is created
	reader := bufio.NewReader(os.Stdin)
	userText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	numberOfWords := len(strings.Fields(userText))
	fmt.Println(numberOfWords)

	// line counter

}
