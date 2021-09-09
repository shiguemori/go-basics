/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	done := make(chan bool, 1)

	go func() {
		msg := "Hello world"
		runes := []rune(msg)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		for i := 0; i < 10000000000; i++ {
		}
		messages <- string(runes)
		done <- true
	}()

	msg := <-messages
	fmt.Println(msg)
	if <-done {
		fmt.Println("Finished")
	}
}
