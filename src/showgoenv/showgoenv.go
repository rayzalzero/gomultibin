package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		user    string
		homeDir string
		gopath  string
	)
	user = os.Getenv("USER")
	homeDir = os.Getenv("HOME")
	gopath = os.Getenv("GOPATH")

	fmt.Println("hallo ", user)
	fmt.Println("home ", homeDir)
	fmt.Println("home ", gopath)
}
