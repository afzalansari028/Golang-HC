package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in golang")

	// myCh := make(chan int) //Unbuffered channel
	myCh := make(chan int, 2) //Buffered channel
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	// RECIEVE ONLY (<-chan)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// SEND ONLY (chan<-)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 6
		myCh <- 5
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
