package main

import (
	"fmt"
	"github.com/startFromBottom/learngo/accounts"
	// "log"
)


func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err) // kill process
		fmt.Println(err)
	}
	fmt.Println(account)
}

