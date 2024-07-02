## Tópicos em Go

## Estruturas, métodos e interfaces

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
 Uma funçao que recebe uma funcção
 ```go
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
 ```

### Defer Panic recover
Defer: escalona um método para que ele seja executado ao fim da função.
 ```go
 package main

import "fmt"

func dia1() {
	fmt.Println("Ola mundo 1")
}
func dia2() {
	fmt.Println("Ola mundo 2")
}

func main() {

	defer dia2()

	dia1()

}
```

Panic: Para a execução com um erro provido pelo programador
Recover: interrompe o Panic

```go

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Panico no trem")
}
```

### Ponteiros e operadores 
Ponteiro é uma variavel que armazena um endereço de memória

&val -> retorna o endereço dessa varialvel
\* -> acessa a variavel naquele endereço 
```go
a := 1 // 1
pont := &a // 0xc000061f60
(*pont) // 1
```
```go
func somaNumero (endNum *int, val int) {
	(*endNum) += val
}
```
## Pacotes
Reutilizar e não reescrever código
Exemplos de pacotes:
- String
- IO
- File/ Filepath

### Strings

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Computador", "com"))
}
```

### OS 

```go 
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
```