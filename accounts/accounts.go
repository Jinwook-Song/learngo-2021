package accounts

import "errors"

// Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("can't withdraw ❌")

// NewAccount creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

// Methods

// Deposit x amount on your account
// * : not copy, pointer receiver
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balacne of your account
func (a Account) Balacne() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} 
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}