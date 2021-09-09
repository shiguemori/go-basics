/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
	for i := 0; i <= 9; i++ {
		str = strings.Replace(str, strconv.FormatInt(int64(i), 10), "", -1)
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return strings.Trim(string(runes), " ")
}

func main() {
	result := ModifyString("oll123eH56")
	fmt.Println(result)
}
