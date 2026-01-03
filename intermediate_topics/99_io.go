package main

import (
	"fmt"
	"io"
	"strings"
)

// 99 IO package
func Demo99IO() {
	fmt.Println("-- 99 IO package --")
	r := strings.NewReader("copy me")
	w := &strings.Builder{}
	io.Copy(w, r)
	fmt.Println("copied:", w.String())
}
