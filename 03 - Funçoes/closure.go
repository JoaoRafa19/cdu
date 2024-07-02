package main


func somaNumero (endNum *int, val int) {
	(*endNum) += val
}

func main() {

	num := 0;

	numero := func() int {
		num ++
		return num;
	}

	println(numero())
	println(num)
}