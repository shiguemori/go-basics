/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_channelBuffering(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "test 1",
			args: []string{"test", "test", "test"},
			want: "testtesttest",
		},
		{
			name: "test 2",
			args: []string{"test2", "test2", "test2"},
			want: "test2test2test2",
		},
		{
			name: "test 3",
			args: []string{"test3", "test3", "test3"},
			want: "test3test3test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			channelBuffering(out, tt.args)
			got := strings.Replace(out.String(), "\n", "", -1)
			compareStringsFunction(got, tt.want, "channelBuffering", t)
		})
	}
}
