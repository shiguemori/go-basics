/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func compareStringsFunction(out, want, funcName string, t *testing.T) {
	if strings.Compare(out, want) != 0 {
		t.Errorf("%v() = %v, want %v", funcName, out, want)
	}
}

func Test_bufferToStringReplace(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test 1",
			args: "ping",
			want: "ping ",
		},
		{
			name: "test 2",
			args: "pong",
			want: "pong ",
		},
		{
			name: "test 3",
			args: "test",
			want: "test ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			fmt.Fprint(out, tt.args+"\n")
			if got := bufferToStringReplace(out); got != tt.want {
				t.Errorf("bufferToStringReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programPointer(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(100))
	wallet.Withdraw(Bitcoin(10))
	wallet.Transfer(Bitcoin(10))

	tests := []struct {
		name string
		want Wallet
	}{
		{
			name: "test 1",
			want: wallet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := programPointer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("programPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programStruct(t *testing.T) {

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
	tests := []struct {
		name string
		want []struct {
			name      string
			shape     Shape
			area      float64
			perimeter float64
		}
	}{
		{
			name: "test 1",
			want: figures,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := programStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("programStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
