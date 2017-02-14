package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	var password string
	flag.StringVar(&password, "password", "DEFAULT", "Password for encryption")
	flag.Parse()

	if password == "" {
		fmt.Println("Password is empty. Still generating")
	}

	p := []byte(password)
	hp, _ := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)

	fmt.Println("\nPassword: ", password)
	fmt.Println("Hashed Password: ", string(hp))
}
