package banking

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("Got: %s Want: %s", got, want)
	}
}