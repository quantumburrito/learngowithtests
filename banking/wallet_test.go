package banking

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	
	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("Got: %d Want: %d", got, want)
	}
}