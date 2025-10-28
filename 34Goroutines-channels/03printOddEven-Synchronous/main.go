package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go PrintOdd(ch1, ch2)
	go PrintEven(ch1, ch2)

	time.Sleep(1 * time.Second)
}

func PrintOdd(ch1, ch2 chan struct{}) {

	for i := 1; i <= 10; i++ {
		<-ch1
		if i%2 != 0 {
			fmt.Println("Odd:", i)
		}
		ch2 <- struct{}{}
	}
}
func PrintEven(ch1, ch2 chan struct{}) {

	for i := 1; i <= 10; i++ {
		ch1 <- struct{}{}
		if i%2 == 0 {
			fmt.Println("Even:", i)
		}
		<-ch2
	}
}
