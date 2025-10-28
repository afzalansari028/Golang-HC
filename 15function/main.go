package main

import "fmt"

func main() {
	fmt.Println("Welcome to the fuction learning in golang")
	greetings()
	result := adder(5, 9)
	fmt.Println("value is :", result)

	proRes, myMessage := proAdder(1, 2, 3, 4)
	fmt.Println("Pro adder values :", proRes, myMessage)

	sl := make([]int, 5)
	sl = append(sl, 1, 2, 3, 4, 5, 6, 7, 8)

	isNonZero := func(n int) bool {
		return n != 0
	}

	newslice := Filter(sl, isNonZero)

	fmt.Printf("len:%d  cap:%d  %v\n", len(sl), cap(sl), sl)
	fmt.Printf("len:%d  cap:%d  %v", len(sl), cap(sl), newslice)
}

func greetings() {
	fmt.Println("Hi Good evening!!")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi pro adder returning"
}

//use function as a type by passing as a parameter
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
