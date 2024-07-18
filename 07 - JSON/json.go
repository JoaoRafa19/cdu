package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)
type users struct {
	Usuarios []User `json:"users"`
}

type Network struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type User struct {
	Name     string    `json:"name"`
	Tel      string    `json:"tel"`
	Networks []Network `json:"networks"`
}

func main() {
	jsonFile, err := os.OpenFile("./file.json", os.O_RDONLY, os.ModeAppend)

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	
	byteValue, _ := io.ReadAll(jsonFile)

	var usuarios users

	if err := json.Unmarshal(byteValue, &usuarios); err != nil {
		panic(err)
	}

	for _, user := range usuarios.Usuarios{
		fmt.Printf("\n\nNome: %s\n", user.Name)
		fmt.Printf("Telefone: %s\n", user.Tel)
		fmt.Printf("Networks: ")
		for _, network := range user.Networks {
			fmt.Printf("\tname: %s\n", network.Name)
			fmt.Printf("\turl: %s\n", network.Url)
		}
	}




}
