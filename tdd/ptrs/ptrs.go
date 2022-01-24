package ptrs

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("insufficient funds, cannot withdraw")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	Balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.Balance += b
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.Balance {
		return ErrInsufficientFunds
	}
	w.Balance -= b
	return nil
}
