package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	fmt.Println("Enter the rating for pizza")
	reader := bufio.NewReader(os.Stdin)

	//comma ok OR comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
