/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	fmt.Println(fmt.Sprintf("Program %d - Is running on http://localhost:5050/YOUR_NAME", 99))
	log.Fatal(http.ListenAndServe(":5050", http.HandlerFunc(MyGreeterHandler)))
}

func bufferToStringReplace(args *bytes.Buffer) string {
	return strings.Replace(args.String(), "\n", " ", -1)
}

func programStruct() []struct {
	name      string
	shape     Shape
	area      float64
	perimeter float64
} {
	figures := []struct {
		name      string
		shape     Shape
		area      float64
		perimeter float64
	}{
		{
			name:      "Rectangle",
			shape:     Rectangle{Width: 12, Height: 6},
			area:      72.0,
			perimeter: 36,
		},
		{
			name:      "Circle",
			shape:     Circle{Radius: 10},
			area:      314.1592653589793,
			perimeter: 62.83185307179586,
		},
		{
			name:      "Triangle",
			shape:     Triangle{Base: 12, Height: 6, Sides: []float64{3.0, 3.0, 3.0}},
			area:      36.0,
			perimeter: 9,
		},
	}
	return figures
}

func programPointer() Wallet {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(100))
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		return wallet
	}
	err = wallet.Transfer(Bitcoin(10))
	if err != nil {
		return wallet
	}
	return wallet
}
