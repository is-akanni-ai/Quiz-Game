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


	score := 0
	numOfQst := 2

	fmt.Print("Qst 1: What is better?, RTX 3080 or RTX 3090 ")

	var answe1 string;
	var answe2 string;

	fmt.Scan(&answe1, &answe2)

	switch answe1 + " " + answe2 {
        case "RTX 3090":
		    fmt.Println("Correct!")
			score += 5
			// numOfQst += 1
			
	    case "rtx 3090":
			fmt.Println("Correct!")
			score += 5
			// numOfQst += 1
	default:
		fmt.Println("Incorrect!")
		// numOfQst += 1
	}

	fmt.Print("Qst 2: How many cores does the AMD Ryzen 9 3900X have? ")
	var cores uint;

	fmt.Scan(&cores);

	if cores == 12 {
		fmt.Println("Correct!")
		score += 5
		// numOfQst += 1
	}else{
		fmt.Println("Incorrect!")
		// numOfQst += 1
	}

	fmt.Println("Qst 3: ")

	fmt.Printf("You score %v out of 20 with %v.\n questions", score, numOfQst);

	percent := (float64(score) / float64(numOfQst)) * 100


	fmt.Printf("You scored:  %v%%", percent);
}