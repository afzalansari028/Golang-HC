package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	receavingval := gen(1, 2, 3, 4, 5, 6, 7, 8)
	even, odd := OddEven(receavingval)

	wg.Add(2)
	go PrintEven(even, wg)
	go PrintOdd(odd, wg)

	wg.Wait()

}

func OddEven(number <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		for nums := range number {
			if nums%2 == 0 {
				evenCh <- nums
			} else {
				oddCh <- nums
			}
		}
		close(evenCh)
		close(oddCh)
	}()
	return evenCh, oddCh
}

func gen(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, val := range nums {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

func PrintEven(even <-chan int, wg *sync.WaitGroup) {
	for e := range even {
		fmt.Println("even---", e)
	}
	wg.Done()
}
func PrintOdd(odd <-chan int, wg *sync.WaitGroup) {
	for o := range odd {
		fmt.Println("odd---", o)
	}
	wg.Done()
}
