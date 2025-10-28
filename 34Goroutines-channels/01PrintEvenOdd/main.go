package main

import "fmt"

func main() {

	evenCh := make(chan int)
	oddCh := make(chan int)

	go PrintEven(evenCh)
	go PrintOdd(oddCh)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			evenCh <- i // sending data to the channel
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func PrintEven(evenCh <-chan int) { // recieving data from the channel
	for num := range evenCh {

		fmt.Println("even number::", num)
	}
}
func PrintOdd(oddCh <-chan int) {
	for num := range oddCh {
		fmt.Println("odd number::", num)
	}
}
