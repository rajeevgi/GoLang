package main

import "fmt"

func main() {

	// Variable declarations
	var name string

	// Take a input from user
	fmt.Print("Enter your name: ");
	fmt.Scanln(&name);

	// Print the result
	fmt.Println("Welcome,", name);
}