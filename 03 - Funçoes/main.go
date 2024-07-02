package main

import "fmt"

func main() {
	lista := []float64{7, 5, 8, 9, 4, 0, 9, 10, 10}
	fmt.Println(media(lista))
}

func media(array [] float64) float64 {
	return total(array) / float64(len(array))
}

func total(array []float64) float64 {
	total := 0.0
	for _, valor := range array {
		total += valor
	}
	return total
}
