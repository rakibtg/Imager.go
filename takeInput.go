package main

import (
	"fmt"
)

// Method to take a string input
func takeStringInput(question string) string {
	fmt.Print("⇒ " + question + " ")
	var text string
	_, err := fmt.Scanf("%s", &text)
	if err != nil {
		fmt.Println("Please provide valid string.")
	}
	return text
}

// Method to take a integer input
func takeIntInput(question string) int {
	fmt.Print("⇒ " + question + " ")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Please provide a valid integer value.")
	}
	return i
}
