package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		fmt.Printf("address of balance in test %v \n", &wallet.balance)

		expected := Bitcoin(20)

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(150)}

		wallet.Withdraw(Bitcoin(100))

		expected := Bitcoin(50)

		assertBalance(t, wallet, expected)
	})

}
