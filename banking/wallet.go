package banking

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

func (w *Wallet) Withdrawl(amount Bitcoin) error  {

	if amount > w.Balance() {
		return errors.New("OH NO!")
	}
	w.balance -= amount
	return nil
}