package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		got := wallet.Balance()

		fmt.Printf("address of balance in test %v \n", &wallet.balance)

		expected := Bitcoin(20)

		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(150)}

		wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()

		expected := Bitcoin(50)

		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

}
