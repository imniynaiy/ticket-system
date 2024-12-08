package main

import (
	"log"

	"github.com/spf13/pflag"
	"golang.org/x/crypto/bcrypt"
)

var (
	saltFlag     = pflag.StringP("salt", "s", "", "hash salt")
	passwordFlag = pflag.StringP("password", "p", "", "password")
	costFlag     = pflag.IntP("cost", "c", 10, "cost (4-31)")
)

func main() {
	pflag.Parse()
	hashByte, err := bcrypt.GenerateFromPassword([]byte(*saltFlag+*passwordFlag), *costFlag)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(hashByte))
	err = bcrypt.CompareHashAndPassword(hashByte, []byte(*saltFlag+*passwordFlag))
	log.Printf("pass test: %v\n", err == nil)
}
