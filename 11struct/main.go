package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// ---for sorting----
type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Less(i int, j int) bool {
	return a[i].Name < a[j].Name
}
func (a ByName) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// --------------------

func main() {
	fmt.Println("Welcome to struct in golang")

	afzal := User{"Afzal", "afzal@gmail.com", true, 23}
	fmt.Println(afzal)
	fmt.Printf("Afzal details are : %+v\n", afzal)
	fmt.Printf("name is %v and email is %v\n", afzal.Name, afzal.Email)

	//sort object(by name)
	people := []Person{
		{Name: "Afzal", Age: 25},
		{Name: "Praveen", Age: 23},
		{Name: "Deepak", Age: 24},
	}

	fmt.Println("before sort people::", people)
	sort.Sort(ByName(people))
	fmt.Println("after sorted by name people::", people)
}
