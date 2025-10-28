package main

import (
	"Golang_H/18Interface-real/bank"
	"fmt"
)

func main() {

	// union := bank.NewUnionBank()
	// union.Deposite(500)
	// b := union.GetBalance()
	// fmt.Println("ubi balance is--", b)
	// union.Withdraw(100)
	// bn := union.GetBalance()
	// fmt.Println("ubi balance is--", bn)

	axis := bank.NewAxisBank()
	axis.Deposite(500)
	axisBalance := axis.GetBalance()
	fmt.Println("axis balance ", axisBalance)

	err := axis.Withdraw(200)
	if err != nil {
		fmt.Println("erorr", err)
		return
	}
	fmt.Println("axis balance:", axis.GetBalance())

}

type Bank interface {
	GetBalance() int
	Deposite(amount int)
	Withdraw(amount int) error
}
