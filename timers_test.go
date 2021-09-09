/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func TestTimers(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
	}{
		{
			name:  "test 1",
			wantW: "Timer 1 firedTimer 2 stopped",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Timers(w)
			out := bufferToStringReplace(w)
			compareStringsFunction(out, tt.wantW, "Timers", t)
		})
	}
}
