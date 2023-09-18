// Test cases for account service

package account

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount("nico")
	if account.Owner() != "nico" {
		t.Error("NewAccount() owner not match")
	}
	if account.Balance() != 0 {
		t.Error("NewAccount() balance not match")
	}
}

func TestDeposit(t *testing.T) {
	account := NewAccount("nico")
	account.Deposit(10)
	if account.Balance() != 10 {
		t.Error("Deposit() balance not match")
	}
}

func TestWithdraw(t *testing.T) {
	account := NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(5)
	if err != nil {
		t.Error("Withdraw() error not match")
	}
	if account.Balance() != 5 {
		t.Error("Withdraw() balance not match")
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	account := NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err == nil {
		t.Error("Withdraw() error not match")
	}
	if account.Balance() != 10 {
		t.Error("Withdraw() balance not match")
	}
}

func TestChangeOwner(t *testing.T) {
	account := NewAccount("nico")
	account.ChangeOwner("lynn")
	if account.Owner() != "lynn" {
		t.Error("ChangeOwner() owner not match")
	}
}

func TestString(t *testing.T) {
	account := NewAccount("nico")
	if account.String() != "nico's account.\nHas: 0" {
		t.Error("String() not match")
	}
}
