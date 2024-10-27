package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		
		if got != want {
			t.Errorf("Wanted: %s, Got: %s", want, got)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("wanted an error but didn't get one!")
		}
	}
	t.Run("Deposit Bitcoin into Wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdrawl Bitcion from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdrawl(10)
		want := Bitcoin(10)
		assertError(t, err)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdrawl insufficient Funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdrawl(Bitcoin(30))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)


	})
	
}