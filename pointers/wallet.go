package pointers

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// (*w).balance is also valid syntax but explicit dereferencing
	// isn't necessary as struct pointers are automatically dereferenced.
	return w.balance
}
