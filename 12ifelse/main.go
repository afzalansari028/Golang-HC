package main

import "fmt"

func main() {

	loginCount := 23
	var result string

	if loginCount > 10 {
		result = "regular user"
	} else {
		result = "not regular user"
	}
	fmt.Println(result)

	//different syntax
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less that 10")
	}

}
