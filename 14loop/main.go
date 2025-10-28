package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop on golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Satuday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//like for each loop

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is : %v and value is : %v\n", index, day)
	// }

	//in place of while loop

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto king
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is :", rougeValue)
		rougeValue++
	}

king:
	fmt.Println("Hey come here")
}
