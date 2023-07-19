package main

import (
	"fmt"

	"github.com/Themath93/learngo/something/mydict"
)

// bank관련
// func main() {
// 	// 은행 계좌를 개설한다고 가정한다. Owner(public) : "mike" 잔액은 0
// 	// 잔액이 입금 출금이 지속된다
// 	// 이때 account 라는 변수의 Owner는 절대로 바뀌면 안된다!! 근데 Owner는 pubic 하다.
// 	// 그래서 바뀐다.
// 	// account := banking.Account{Owner: "mike"}
// 	// fmt.Println(account)
// 	// {mike 0}
// 	// account.Owner = "pepe"
// 	// fmt.Println(account)
// 	// {pepe 0}

// 	account := accounts.NewAccount("mike")
// 	// fmt.Println(account)
// 	// account.owner 는 접근할 수 없는 private 하기 때문에 account의 owner 는 변경 불가능해진다.

// 	// Deposit
// 	account.Deposit(10)

// 	// Withdraw
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(account.Balance())

// 	// account.ChangeOwner("paul")
// 	// fmt.Println(account.Owner(), account.Balance())

// 	fmt.Println(account)
// }

// dict

func main() {
	dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"

	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found:", word, "definition:", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// // hello2, _ := dictionary.Search(word)

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
