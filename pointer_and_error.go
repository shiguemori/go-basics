/*
 * Copyright (c) 2021.
 */

package main

import (
	"errors"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Transfer(amount Bitcoin) error {
	w.Withdraw(amount)
	w.balance -= 3
	return nil
}
