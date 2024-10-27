package banking

import "fmt"

type Bitcoin int

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	fmt.Printf("Address of the balance in wallet is:%p \n", &w.balance)
	return (*w).balance
}