package main

import "fmt"

func ifElseExample() {
	var a = 10
	var b = 2
	result := a / b
	if result >= 5 {
		fmt.Println(true)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// same as nodejs. can use if, or if else or if else if else. but the parenthesis is important
