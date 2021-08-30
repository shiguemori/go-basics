/*
 * Copyright (c) 2021.
 */

package main

import "fmt"

func main() {
	fmt.Println(fmt.Sprintf("Program %d - Hello(\"Shiguemori\", \"\") return %s", 1, Hello("Shiguemori", "")))
	fmt.Println(fmt.Sprintf("Program %d - Add([]int{5, 4, 9, 1, 3, 6}) return %d", 2, Add([]int{5, 4, 9, 1, 3, 6})))
	fmt.Println(fmt.Sprintf("Program %d - Sub([]int{5, 4, 9, 1, 3, 6}) return %d", 2, Sub([]int{5, 4, 9, 1, 3, 6})))
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(100))
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		return
	}
	err = wallet.Transfer(Bitcoin(10))
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Program %d - wallet.Balance() return %d", 3, wallet.Balance()))

}
