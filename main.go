package main

import (
	"fmt"
	// "os"
	// "bufio"
	// "strings"
)

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
	numOfQst := 4
	qstAnswered := 0

	fmt.Print("Qst 1: What is better?, RTX 3080 or RTX 3090 ")

	var answe1 string;
	var answe2 string;

	fmt.Scan(&answe1, &answe2)

	switch answe1 + " " + answe2 {
        case "RTX 3090":
		    fmt.Println("Correct!")
			score ++			
	    case "rtx 3090":
			fmt.Println("Correct!")
			score ++
	default:
		fmt.Println("Incorrect!")
	}
	qstAnswered++


	fmt.Print("Qst 2: How many cores does the AMD Ryzen 9 3900X have? ")
	var cores uint;

	fmt.Scan(&cores);

	if cores == 12 {
		fmt.Println("Correct!")
		score ++
	}else{
		fmt.Println("Incorrect!")
	}
	qstAnswered++


	fmt.Print("Qst 3: What does CPU stand for? ")

	var cpuAns1 string;
	var cpuAns2 string;
	var cpuAns3 string;

	fmt.Scan(&cpuAns1, &cpuAns2, &cpuAns3);

	switch cpuAns1 + " " + cpuAns2 + " " + cpuAns3 {
	    case "Central Processing Unit":
		    fmt.Println("Correct!")
		    score ++
	    case "central processing unit":
		    fmt.Println("Correct!")
		    score ++
	    default: 
		    fmt.Println("Incorrect!")
	}

	qstAnswered ++


	fmt.Printf("Qst 4: What is the capital of France? ");

	var capital string;
	fmt.Scan(&capital);

    if capital == "Paris" || capital == "paris" {
		fmt.Println("Correct!")
		score++
		qstAnswered += 1
	} else {
		fmt.Println("Incorrect!")
		qstAnswered += 1
	}


	// fmt.Print("Qst 4: What is the capital of France? ")
    // capital, _ := reader.ReadString('\n')
    // capital = strings.TrimSpace(capital)

    //  if capital == "Paris" || capital == "paris" {
    //     fmt.Println("Correct!")
    //     score++
    // } else {
    //     fmt.Println("Incorrect!")
    // }
    // qstAnswered++


	fmt.Printf("Congrats %v, you have completed the quiz!\n", name)
	fmt.Printf("You anwered %v correctly out of %v questions, total questions answered  %v. questions \n", score, numOfQst, qstAnswered);

	percent := (float64(score) / float64(numOfQst)) * 100

	fmt.Printf("You scored:  %.0f%%\n", percent)
}