package accounts

import (
	"fmt"
	"errors"
)

// Account Struct	
type Account struct { // if lowercase: private, uppercase: public
	owner string 
	balance	int
}

var errNoMoney = errors.New("Can't withdraw. you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account	
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// a : receiver (약속 : first lowercase of struct)
	// * : Don't make copy of Account (없으면 복제본에 값 더해짐...)
	fmt.Println("Gonna Deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
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

// Owner of an account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	// like python __str__ method
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance() )
}