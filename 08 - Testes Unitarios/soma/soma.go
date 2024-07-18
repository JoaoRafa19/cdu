package soma



func Soma(a int, b int) int {
	return a + b
}

func Fat(a int) int {
	if a == 0 {
		return 1
	}
	return a * Fat(a -1)
}

func SomaNums(i ...int) int {

	var val int 

	for _, v := range i {
		val += v
	}

	return val
}