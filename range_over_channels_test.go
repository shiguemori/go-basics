/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func TestRangeOverChannels(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
	}{
		{
			name:  "test 1",
			wantW: "one two ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			RangeOverChannels(w)
			out := bufferToStringReplace(w)
			compareStringsFunction(out, tt.wantW, "RangeOverChannels", t)
		})
	}
}
