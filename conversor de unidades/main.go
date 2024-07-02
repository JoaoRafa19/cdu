package main

import "fmt"

func main() {

	for {
		var entrada int
		
		fmt.Println("Digite a entrada: (-1 para sair)")
		_, err := fmt.Scanf("%d ", &entrada)
		if err != nil {
			continue
		}
		if entrada != -1 {
			c := entrada - 273
			fmt.Printf("%d K = %d ºC\n", entrada, c)

		} else {
			break
		}
	}

}
