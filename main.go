package main

import (
	"flag"
	"fmt"

	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	var password = flag.String("password", "", "Password for encryption")

	if *password != "" {
		log.Println("Password is empty. Still generating")
	}

	p := []byte(*password)
	hp, _ := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)

	fmt.Println("\nHashed Password: ", string(hp))
}
