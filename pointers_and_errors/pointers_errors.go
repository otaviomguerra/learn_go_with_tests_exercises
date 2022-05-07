package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//String denotes the behavior when wanting to print a Bitcoin type.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

//Deposit deposits a given amount into a Bitcoin Wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance gives the current balance of a given Bitcoin Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw withdraws a given Bitcoin amount from the Bitcoin Wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
