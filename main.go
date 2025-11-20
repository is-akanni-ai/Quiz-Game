package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")

	fmt.Printf("Enter your name: ")

	var name string;
	fmt.Scan(&name)
	fmt.Printf("Hello, %s! Let's start the quiz.\n", name)
	
	fmt.Println("Enter your age: ")

	var age uint;
	fmt.Scan(&age)
	fmt.Println(age >= 10)
}