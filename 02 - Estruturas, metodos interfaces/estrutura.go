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
