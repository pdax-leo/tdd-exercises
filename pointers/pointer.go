package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin struct
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit func
func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount

}

// Balance func
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw func
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
