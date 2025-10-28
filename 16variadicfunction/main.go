package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4}
	add(vals...)
	// vals = []int{1, 2, 3, 4, 5, 6}
	// add(vals...)

	add(1, 2, 3, 4)
	add(1, 2, 3, 4, 5, 6)

	printAll(1, 23, "Hello")
}

//  pass a slice of interface{}
func printAll(arg ...interface{}) {
	fmt.Println(arg...) // forward it here
}

func add(vals ...int) {
	// fmt.Println("vals--", vals)
	total := 0
	for _, val := range vals {
		total += val
	}
	fmt.Println(vals, total)
}
