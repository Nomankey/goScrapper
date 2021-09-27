package main

import (
	"fmt"
	"log"

	"github.com/nomankey/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nomankey")

	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	account.Withdraw(5)
	fmt.Println(account.Balance())
}