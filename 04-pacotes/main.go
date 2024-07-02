package main

import (
	"bytes"
	"fmt"
)

// func main() {
// 	if _, err := io.WriteString(os.Stdout, "Hello world"); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	message := []byte("hello, gophers!")

	// err := os.WriteFile("hello", message, os.ModeAppend.Perm())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%s", bytes.Title(message))

}
