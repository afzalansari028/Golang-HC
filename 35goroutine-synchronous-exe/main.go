package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	//   print foo after bar (synchronous execution using goroutine and channel)
	go Foo(ch1, ch2)
	go Bar(ch1, ch2)

	time.Sleep(1 * time.Second)
}

func Bar(ch1, ch2 chan struct{}) {
	for i := 0; i < 3; i++ {
		ch1 <- struct{}{} //wait for foo
		<-ch2             // print bar
		fmt.Println("Bar")
	}
}

func Foo(ch1, ch2 chan struct{}) {
	for i := 0; i < 3; i++ {
		<-ch1 //print foo
		fmt.Println("Foo")
		ch2 <- struct{}{} //wait for bar
	}
}
