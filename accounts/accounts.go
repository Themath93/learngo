package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
	// account 라는 값의 메모리주소(pointer)를 리턴하고
	// 타입은 *Account 로 리턴한다.
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// 여기서 a 는 receiver 라고 하는데 Account 타입을 상속받아 사용하는 것 같다.
	a.balance += amount

}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOnwer of the account
func (a *Account) ChangeOwner(newOnwer string) {
	a.owner = newOnwer
}

// Onwer of the acoount
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
