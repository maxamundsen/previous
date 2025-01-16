package main

import (
	"hotlap/crypt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		passHash, _ := crypt.HashPassword(os.Args[1])
		println(passHash)
	} else {
		println("Please input a password as first program argument")
	}
}