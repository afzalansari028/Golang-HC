package main

import "fmt"

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Tom Tompson",
			Id:   "1231231",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.Id)
	fmt.Println(m.Description())
}

type Employee struct {
	Name string
	Id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.Id)
}
func Description() string {
	var e Employee
	return fmt.Sprintf("%s (%s)", e.Name, e.Id)
}

type Manager struct {
	Employee
	Reports []Employee
}

// func (m Manager) FindNewEmployees() []Employee {
// 	//business loginc
// }
