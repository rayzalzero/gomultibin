package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		user    string
		homeDir string
	)
	user = os.Getenv("USER")
	homeDir = os.Getenv("HOME")

	fmt.Println("hallo ", user)
	fmt.Println("home ", homeDir)
}
