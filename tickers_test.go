/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func TestTickers(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
	}{
		{
			name:  "test 1",
			wantW: "Ticker stopped ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Tickers(w)
			got := bufferToStringReplace(w)
			compareStringsFunction(got, tt.wantW, "Tickers", t)
		})
	}
}
