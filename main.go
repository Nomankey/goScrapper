package main

import (
	"fmt"

	"github.com/nomankey/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nomankey")
	fmt.Println(account)
}