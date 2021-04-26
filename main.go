package main

import (
	"fmt"

	"github.com/jinwook-song/learngo-2021/accounts"
)

func main() {
	accounts := accounts.NewAccount("nico")
	accounts.Deposit(10)
	fmt.Println(accounts.Balacne())
	// error handle
	err := accounts.Withdraw(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accounts )
	accounts.ChangeOwner("jinwook")
	fmt.Println("changed owner:" ,accounts.Owner())
}