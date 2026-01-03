package main

import (
	"fmt"
	"os"
)

// 88 Temporary Files and Directories
func Demo88TempFilesDirs() {
	fmt.Println("-- 88 Temporary Files and Directories --")
	f, err := os.CreateTemp("", "gotutfile-*.txt")
	if err != nil {
		fmt.Println("create temp file error:", err)
		return
	}
	name := f.Name()
	f.WriteString("temp content\n")
	f.Close()
	fmt.Println("Created temp file:", name)
	_ = os.Remove(name)
}
