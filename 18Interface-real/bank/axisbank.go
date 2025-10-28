package bank

import "errors"

type AxisBank struct {
	Balance int
	Changes int
}

func NewAxisBank() *AxisBank {
	return &AxisBank{
		Balance: 0,
		Changes: 10,
	}
}
func (ax *AxisBank) GetBalance() int {
	return ax.Balance
}

func (ax *AxisBank) Deposite(amount int) {
	ax.Balance += amount
}

func (ax *AxisBank) Withdraw(amount int) error {
	newBalance := ax.Balance - amount - ax.Changes

	if newBalance < 0 {
		return errors.New("Insufficiant balance!!")
	}
	ax.Balance = newBalance
	return nil
}
