package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	// myDefer()
	ValidateAge(12)
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func ValidateAge(age int) {
	fmt.Println("hii")
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered panic:", r)
		}
	}()

	if age < 18 {
		panic("You are below 18, you can't vote...")
	} else {
		fmt.Println("You are eligible for vote")
	}
	fmt.Println("byeee")
}
