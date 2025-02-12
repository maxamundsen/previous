package main

import (
	"os"
	"previous/security"
)

func main() {
	if len(os.Args) == 2 {
		passHash, _ := security.HashPassword(os.Args[1])
		println(passHash)
	} else {
		println("Please input a password as first program argument")
	}
}
