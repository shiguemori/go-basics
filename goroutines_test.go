/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_runGoRoutine(t *testing.T) {
	t.Run("running runGoRoutine", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		runGoRoutine(buffer)
		got := strings.Replace(buffer.String(), "\n", " ", -1)
		want := "direct:0 direct:1 direct:2 going goroutine:0 goroutine:1 goroutine:2 done "
		if strings.Compare(got, want) != 0 {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func Test_routineLoop(t *testing.T) {
	type args struct {
		io.Writer
		from string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "Test 1",
			args: args{
				&bytes.Buffer{},
				"shiguemori",
			},
			wantOut: "shiguemori:0 shiguemori:1 shiguemori:2 ",
		},
		{
			name: "Test 2",
			args: args{
				&bytes.Buffer{},
				"vinicius",
			},
			wantOut: "vinicius:0 vinicius:1 vinicius:2 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			routineLoop(out, tt.args.from)
			gotOut := strings.Replace(out.String(), "\n", " ", -1)
			if strings.Compare(gotOut, tt.wantOut) != 0 {
				t.Errorf("routineLoop() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
