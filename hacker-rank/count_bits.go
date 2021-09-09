/*
 * Copyright (c) 2021.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) (bits int32) {
	bits = int32(strings.Count(fmt.Sprintf("%b", num), "1"))
	return
}

func main() {
	bits := countBits(127)
	fmt.Println(bits)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
