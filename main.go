package main

import (
	"fmt"

	"github.com/nomankey/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nomankey")
	account.Deposit(10)
	fmt.Println(account)
}