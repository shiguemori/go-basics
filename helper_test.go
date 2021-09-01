/*
 * Copyright (c) 2021.
 */

package main

import (
	"strings"
	"testing"
)

func compareStringsFunction(out, want, funcName string, t *testing.T) {
	if strings.Compare(out, want) != 0 {
		t.Errorf("%v() = %v, want %v", funcName, out, want)
	}
}
