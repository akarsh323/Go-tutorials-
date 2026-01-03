package main

import "fmt"

import "os"

// 92 Environment Variables
func Demo92EnvVars() {
	fmt.Println("-- 92 Environment Variables --")
	fmt.Println("SHELL:", os.Getenv("SHELL"))
}
