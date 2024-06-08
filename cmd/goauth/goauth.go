package main

import (
	"crypto/sha256"
	"fmt"
	"goauth/internal/User"
)

func main() {
	fmt.Println("I live!")

	pw := "P@ssw0rd"

	rsh := User.CreateUser("Ryan Shaw-Harrison", "rsh", pw, "rsh@mail.local", User.StandardUser, true)

	fmt.Printf("%s: %s -> %x\n", rsh.GetUsername(), rsh.StringifyUsertype(), rsh.GetPassword())

	fmt.Printf("%x\n", hashPassword(pw))

	fmt.Printf("%+v\n", rsh)
}

func hashPassword(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hash.Sum(nil)
}
