package main

import "fmt"

func forTest() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	j := 3
	for j <= 3 {
		fmt.Println(j)
		break
	}

	const t = 5
	for t == 5 {
		fmt.Println(t)
		break
	}

	for i := range 2 {
		fmt.Println("range ", i) // iterates over the a sequence. it will print 0,1
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
