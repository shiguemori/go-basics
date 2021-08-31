/*
 * Copyright (c) 2021.
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, strings.Replace(r.URL.Path, "/", "", -1))
}
