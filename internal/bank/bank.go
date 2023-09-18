/*
Application that simulates bank functionality. Using account app creates accounts, and stores logs of opperations in a file.
reads transaction queue from a file and executes them.
*/
package bank

import (
	"fmt"
	"gobank/internal/account"
)

// Bank struct
type Bank struct {
	accounts []*account.Account
	transactions []string
}

// NewBank creates Bank
func NewBank() *Bank {
	bank := Bank{}
	return &bank
}

// CreateAccount creates new account
func (b *Bank) CreateAccount(owner string) *account.Account {
	account := account.NewAccount(owner)
	b.accounts = append(b.accounts, account)
	return account
}

// Deposit money to account
func (b *Bank) Deposit(account *account.Account, amount int) {
	account.Deposit(amount)
	b.transactions = append(b.transactions, fmt.Sprintf("%s %d", "deposit", amount))
}

// Withdraw money from account
func (b *Bank) Withdraw(account *account.Account, amount int) error {
	err := account.Withdraw(amount)
	if err != nil {
		return err
	}
	b.transactions = append(b.transactions, fmt.Sprintf("%s %d", "withdraw", amount))
	return nil
}

// Transfer money from one account to another
func (b *Bank) Transfer(from *account.Account, to *account.Account, amount int) error {
	err := from.Withdraw(amount)
	if err != nil {
		return err
	}
	to.Deposit(amount)
	b.transactions = append(b.transactions, fmt.Sprintf("%s %d %s %s", "transfer", amount, from.Owner(), to.Owner()))
	return nil
}

// ChangeOwner of the account
func (b *Bank) ChangeOwner(account *account.Account, newOwner string) {
	account.ChangeOwner(newOwner)
}

// Owner of the account
func (b *Bank) Owner(account *account.Account) string {
	return account.Owner()
}