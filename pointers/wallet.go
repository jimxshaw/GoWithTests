package pointers

import (
	"errors"
	"fmt"
)

// Creating a new type from existing one:
// type MyName OriginalType
// This is useful as methods can be
// declared on the new type.
type Bitcoin int

// Implementing Stringer interface on new Bitcoin type.
// This defines how your type is print when using %s format.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// (*w).balance is also valid syntax but explicit dereferencing
	// isn't necessary as struct pointers are automatically dereferenced.
	return w.balance
}
