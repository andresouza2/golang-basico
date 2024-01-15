package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "Fusca 2"
	fmt.Println(c.Name, "andou")
}

func main() {

	carro := Carro{
		Name: "Fusca",
	}

	carro.andou()


	// memoria-endereço(10) <-- a <-- 10

	// a := 10
	// var ponteiro *int = &a
	// fmt.Println(ponteiro) // imprimo o endereço de memoria
	// fmt.Println(*ponteiro) // imprimo o valor que está no endereço de memoria
	// *ponteiro = 50 // altero o valor que está no endereço de memoria
	// fmt.Println(a)
	// fmt.Println(*ponteiro)

	// variavel := 10
	// abc(&variavel)
	// fmt.Println(variavel)

}

// func abc(a *int) {
// 	*a = 500
// }
