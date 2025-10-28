package bank

import "errors"

type Unionbank struct {
	Balance int
}

func NewUnionBank() *Unionbank {
	return &Unionbank{
		Balance: 0,
	}
}
func (ubi *Unionbank) GetBalance() int {
	return ubi.Balance
}

func (ubi *Unionbank) Deposite(amount int) {
	ubi.Balance += amount
}

func (ubi *Unionbank) Withdraw(amount int) error {
	newBalance := ubi.Balance - amount

	if newBalance < 0 {
		return errors.New("Insufficiant balance!!")
	}
	ubi.Balance = newBalance
	return nil
}
