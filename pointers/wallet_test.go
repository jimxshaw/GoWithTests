package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(20)

	got := wallet.Balance()
	expected := 20

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}

}
