## Tópicos em Go

## 02 Estruturas, métodos e interfaces

Exemplo de estrutura e método

```go
package main

type Person struct {
	nome  string
	idade int
}

type Retangulo struct {
	largura, altura float32
}

func (r *Retangulo) area() float32 {
	return r.altura * r.largura
}

```
 Main.go
```go
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
```

## Funções 
```go
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
```

#### Closure