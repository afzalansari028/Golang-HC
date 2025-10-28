package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome the math class in golang")

	// var myNumberOne int = 4
	// var myNumberTwo float64 = 4.5
	// fmt.Println("Sum of two number is :", myNumberOne+int(myNumberTwo))

	// random number

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)

	// random number from crypto

	// myRandom, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(myRandom)

}
