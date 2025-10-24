package main

import "fmt"

func main() {

	// Control Structures : If-else , for-loop and switch case.

	// 1. If-else

	var age int;

	fmt.Print("Enter age: ");
	fmt.Scanln(&age);

	if age > 18 {
		fmt.Println("Adult");
	}else{
		fmt.Println("Minor");
	}

	// 2. for-loop

	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	// 3. Switch case

	var days string = "Monday";

	switch days {

	case "Monday" : 
		fmt.Println("Start of week");
	
	case "Friday" :
		fmt.Println("Weekend soon!");

	default :
		fmt.Println("Midweek...");
	}

}