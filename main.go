package main

import (
	"fmt"
)

type Carro struct {
	Nome string
}

func (c Carro) Andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{Nome: "Gol"}

	carro.Andar()

	resultado := func(x ...int) func() int {

		res := 0

		for _,v := range x {
			res += v
		}
		return func() int {
			return res * res
		}

	} 

	// resultado := somaTudo(10, 5, 8, 20, 154, 1052)

	fmt.Println(resultado(54, 54, 54, 54)())

}

func soma(x int, y int) (result int) {
	result = x + y
	return
}

func somaTudo(x ...int) int {
	resultado := 0

	for _,v := range x {
		resultado += v
	}
	return resultado
}
