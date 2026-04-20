package main

import "fmt"

func varl() {
	fmt.Println("Variables")

	var a string = "Test"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c int = 3 + 3
	fmt.Println(c)

	var e bool = false
	fmt.Println(e)

	f := "short declaration"
	fmt.Println(f)

	g := 2
	fmt.Println(g)
}

// := - using this go will figure for you (only for var - value that can change)
// you can user var <name> <type> too
