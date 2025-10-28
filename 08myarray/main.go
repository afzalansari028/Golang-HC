package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is : ", fruitList)
	fmt.Println("Fruit List is : ", len(fruitList))
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	var vegList = [5]string{"Potato", "beans", "mushroom"}
	fmt.Println("Vegy list is :", len(vegList))

	// a := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	// var b [10]int
	// fmt.Println(len(b))

}
