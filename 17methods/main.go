package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")
	//no inheritance in golang; No super parent

	afzal := User{"Afzal", "afzal@gmail.com", true, 23}
	fmt.Println(afzal)
	fmt.Printf("Afzal details are : %+v\n", afzal)
	fmt.Printf("Afzal name is %v and email is %v\n", afzal.Name, afzal.Email)

	afzal.GetStatus()
	afzal.NewMail()
	fmt.Printf("Afzal name is %v and email is %v\n", afzal.Name, afzal.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("User new mail is :", u.Email)
}
