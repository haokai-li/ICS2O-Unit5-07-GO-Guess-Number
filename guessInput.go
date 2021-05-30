// Created by: haokai
// Created on: May 2021
// This program displays, "While-Loops"

package main

import (
	"fmt"

	"math/rand"

	"time"
)

func main() {

	// This function does addition
	var number int
	var guess int

	rand.Seed(time.Now().UnixNano())

	number = rand.Intn(10)

	fmt.Println("This program is Guess Number.")

	for number != guess{
		// input
		fmt.Println()
		fmt.Print("Enter first number: ")
		fmt.Scanln(&guess)

		if number < guess {
			fmt.Println("It is high!")
		} else if number > guess {
			fmt.Println("It is low!")
		} else {
			fmt.Println("You are right!")
		}
	}

	fmt.Println("\nDone.")
}
