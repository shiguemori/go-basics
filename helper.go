/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"strings"
)

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
