package main

import (
	"errors"
	"fmt"
)

// fake db
var users []User = []User{
	{Id: 10, Name: "Hardik"},
	{Id: 20, Name: "Virat"},
	{Id: 30, Name: "Sanju"},
}

// interface
type UserRepository interface {
	GetById(id int) (User, error)
}

// user type
type User struct {
	Id   int
	Name string
}

type UserService struct {
	repo UserRepository
}

// implementation
type fakeRepository struct{}

func (f *fakeRepository) GetById(id int) (User, error) {

	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return User{}, errors.New("User not found in db")
}

// constructor - initialize user service
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// business logic function
func (service *UserService) GetUser(id int) (User, error) {

	user, err := service.repo.GetById(id)
	if err != nil {
		fmt.Println("Errr occured while fetching")
		return User{}, errors.New("Errr occured while fetching")
	}
	return user, nil
}

// func main() {

// 	fmt.Println("users", users)

// 	fakeRepo := &fakeRepository{}
// 	service := NewUserService(fakeRepo)

// 	user, err := service.GetUser(10)
// 	if err != nil {
// 		fmt.Println("err::", err)
// 		return
// 	}
// 	fmt.Println("user::::::", user)

// }
