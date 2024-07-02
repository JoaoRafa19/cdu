package main

import "fmt"

func dia1() {
	fmt.Println("Ola mundo 1")
}
func dia2() {
	fmt.Println("Ola mundo 2")
}

func defere() {

	defer dia2()

	dia1()

}

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
	}()

	panic("Panico no trem")

}
