package main

import (
	"fmt"
	"os"
)

// 87 Directories
func Demo87Directories() {
	fmt.Println("-- 87 Directories --")
	d, err := os.MkdirTemp("", "gotutdirs")
	if err != nil {
		fmt.Println("mkdirtemp error:", err)
		return
	}
	defer os.RemoveAll(d)
	fmt.Println("Created temp dir:", d)
}
