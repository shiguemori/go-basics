/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func Test_selectFunction1(t *testing.T) {
	tests := []struct {
		name     string
		message1 string
		message2 string
		wantOut1 string
		wantOut2 string
	}{
		{
			name:     "test 1",
			message1: "msg1",
			message2: "msg2",
			wantOut1: "received msg1",
			wantOut2: "received msg2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out1 := &bytes.Buffer{}
			out2 := &bytes.Buffer{}
			selectFunction(out1, out2, tt.message1, tt.message2)
			compareStringsFunction(out1.String(), tt.wantOut1, "selectFunction", t)
			compareStringsFunction(out2.String(), tt.wantOut2, "selectFunction", t)
		})
	}
}
