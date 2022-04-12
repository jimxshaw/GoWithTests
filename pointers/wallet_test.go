package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(20)

	got := wallet.Balance()

	fmt.Printf("address of balance in test %v \n", &wallet.balance)

	expected := 20

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}

}
