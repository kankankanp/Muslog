package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedPassword := "$2a$10$O26qV4vNDJ.79TBkcImRZuJmsljgidQBo8FoJP39djNJloP5O9hQ."
	plaintextPassword := "password"

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if err != nil {
		fmt.Println("Comparison failed:", err)
	} else {
		fmt.Println("Comparison successful!")
	}
}
