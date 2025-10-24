package main

import "fmt"

func main() {
	fmt.Println("Hello from golang , rajeev....");

	// Variables , Constants and Types Example
	var name string = "Rajeev";    // Variable declarations method
	city := "Navi Mumbai";   // Shorthand method for variable declarations
	const pi = 3.143;

	fmt.Printf("Hello %s from %s, the value of pi is : %.2f\n ", name, city, pi);

}