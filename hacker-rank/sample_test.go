/*
 * Copyright (c) 2021.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fizzBuzz(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
			break
		case i%3 == 0:
			fmt.Println("Fizz")
			break
		case i%5 == 0:
			fmt.Println("Buzz")
			break
		default:
			fmt.Println(i)
			break
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fizzBuzz(n)
}
