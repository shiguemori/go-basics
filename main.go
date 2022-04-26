/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(fmt.Sprintf("Program %d - Hello(\"Shiguemori\", \"\") return %s", 1, Hello("Shiguemori", "")))
	fmt.Println(fmt.Sprintf("Program %d.1 - Add([]int{5, 4, 9, 1, 3, 6}) return %d", 2, Add([]int{5, 4, 9, 1, 3, 6})))
	fmt.Println(fmt.Sprintf("Program %d.2 - Sub([]int{5, 4, 9, 1, 3, 6}) return %d", 2, Sub([]int{5, 4, 9, 1, 3, 6})))
	for key, figure := range programStruct() {
		fmt.Println(fmt.Sprintf("Program %d.%d - %s.shape.Area() return %f", 3, key+1, figure.name, figure.shape.Area()))
		fmt.Println(fmt.Sprintf("Program %d.%d - %s.shape.Perimeter() return %f", 3, key+1, figure.name, figure.shape.Perimeter()))
	}
	wallet := programPointer()
	fmt.Println(fmt.Sprintf("Program %d - wallet.Balance() return %d", 4, wallet.Balance()))

	dictionary := Dictionary{"test": "this is just a test", "teste": "Isso e um teste"}
	search, err := dictionary.Search("test")
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Program %d - dictionary.Search(\"test\") return %s", 5, search))

	out := &bytes.Buffer{}
	runGoRoutine(out)
	fmt.Println(fmt.Sprintf("Program %d - runGoRoutine(out) return %s", 6, bufferToStringReplace(out)))

	out2 := &bytes.Buffer{}
	channels(out2, "ping")
	fmt.Println(fmt.Sprintf("Program %d - channels(out2, \"ping\") return %s", 7, bufferToStringReplace(out2)))

	out3 := &bytes.Buffer{}
	channelBuffering(out3, []string{"test", "test", "test"})
	fmt.Println(fmt.Sprintf("Program %d - channelBuffering(out3, []string{\"test\", \"test\", \"test\"}) return %s", 8, bufferToStringReplace(out3)))

	out4 := &bytes.Buffer{}
	channelSynchronization(out4)
	fmt.Println(fmt.Sprintf("Program %d - channelSynchronization(out4) return %s", 9, bufferToStringReplace(out4)))

	out5 := &bytes.Buffer{}
	channelDirections(out5)
	fmt.Println(fmt.Sprintf("Program %d - channelDirections(out5) return %s", 10, bufferToStringReplace(out5)))

	out6 := &bytes.Buffer{}
	out7 := &bytes.Buffer{}
	selectFunction(out6, out7, "test", "test 2")
	fmt.Println(fmt.Sprintf("Program %d - selectFunction(out6, out7, \"test\", \"test 2\") return %s", 11, bufferToStringReplace(out6)+bufferToStringReplace(out7)))

	out8 := &bytes.Buffer{}
	NonBlockingChannel(out8)
	fmt.Println(fmt.Sprintf("Program %d - NonBlockingChannel(out8) return %s", 12, bufferToStringReplace(out8)))

	out9 := &bytes.Buffer{}
	ClosingChannels(out9)
	fmt.Println(fmt.Sprintf("Program %d - ClosingChannels(out9) return %s", 13, bufferToStringReplace(out9)))

	out10 := &bytes.Buffer{}
	RangeOverChannels(out10)
	fmt.Println(fmt.Sprintf("Program %d - RangeOverChannels(out10) return %s", 14, bufferToStringReplace(out10)))

	out11 := &bytes.Buffer{}
	Timers(out11)
	fmt.Println(fmt.Sprintf("Program %d - Timers(out11) return %s", 15, bufferToStringReplace(out11)))

	out12 := &bytes.Buffer{}
	Tickers(out12)
	fmt.Println(fmt.Sprintf("Program %d - Tickers(out12) return %s", 16, bufferToStringReplace(out12)))

	out13 := &bytes.Buffer{}
	Generics(out13)
	fmt.Println(fmt.Sprintf("Program %d - Generics(out13) return %s", 17, bufferToStringReplace(out13)))

	fmt.Println(fmt.Sprintf("Program %d - Is running on http://localhost:5050/YOUR_NAME", 99))
	log.Fatal(http.ListenAndServe(":5050", http.HandlerFunc(MyGreeterHandler)))
}
