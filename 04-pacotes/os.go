package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_CREATE, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte("Ola mundo da alsçdkjfaçlsdjf"))

}