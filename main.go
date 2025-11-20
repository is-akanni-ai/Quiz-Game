package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")

	fmt.Printf("Enter your name: ")

	var name string;
	fmt.Scan(&name)
	fmt.Printf("Hello, %s! Let's start the quiz.\n", name)
	
	fmt.Printf("Enter your age: ")

	var age uint;
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	}else{
		fmt.Printf("Yay ypu can not play!")
	}

	fmt.Print("Qst 1: What is better?, RTX 3080 or RTX 3090 ")

	var answe1 string;
	var answe2 string;

	fmt.Scan(&answe1, &answe2)

	switch answe1 + " " + answe2 {
        case "RTX 3090":
		    fmt.Println("Correct!")
	    case "rtx 3090":
			fmt.Println("Correct!")
	default:
		fmt.Println("Incorrect!")
	}

	fmt.Print("Qst 2: ")
	
}