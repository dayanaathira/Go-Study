package main

import (
	"fmt"
	"time"
)

type Person interface {
	WhoAmI()
}

type Human struct{} // Go method of DTO
type Ghost struct{} // Go method of DTO

func (h Human) WhoAmI() {
	fmt.Println("I am a human")
}
func (g Ghost) WhoAmI() {
	fmt.Println("I am a ghost")
}

func amIA(p Person) {
	switch k := p.(type) {
	case Human:
		k.WhoAmI()
	case Ghost:
		k.WhoAmI()
	default:
		fmt.Println("unknown")
	}
}

func switchExample() {

	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/*
		this like a OOP/nestjs style of :
		class Human {
			whoAmI() { console.log("I am a human") }
		}
		const h = new Human()
		h.whoAmI()
	*/

	amIA(Human{})
	amIA(Ghost{})
}

// intrface is a rulebok. eg: type Animal interface { speak() move() };
// if cat is only move() but dont speak then it is not a Animal.
// if dog are speak() and move() then it is a Animal.
