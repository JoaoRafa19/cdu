package main

import "fmt"

func main() {
	var person = Person{nome: "Jose da silva", idade: 45}
	fmt.Println("Hello World");
	fmt.Println(person);

	var altura float32;
	var largura float32;
	fmt.Scanf("%f %f", &altura, &largura);
	var ret = Retangulo{altura: altura, largura: largura};

	fmt.Println(ret);
	fmt.Println("A area :", ret.area());
}