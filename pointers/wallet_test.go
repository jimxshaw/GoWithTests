package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
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

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s expect %s", got, expected)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but expected one")
	}

	if got != expected {
		t.Errorf("got %q expect %q", got, expected)
	}
}
