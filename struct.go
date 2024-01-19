package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) ola() {
	fmt.Println("Olá, meu nome é", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func main() {

	cliente := Cliente{
		Nome:  "André",
		Email: "andre@gmail.com",
		CPF:   12345678900,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Andréa", "andrea@hotmail.com", 98765432100 }

	fmt.Printf("nome: %s, email: %s, CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome: "David",
			Email: "david@email.com",
			CPF: 12345678900,
		},
		Pais: "EUA", 
	}


	fmt.Printf("nome: %s, email: %s, CPF: %d, País: %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

	cliente.ola()
	cliente2.ola()
	cliente3.ola()

	clienteJson, err := json.Marshal(cliente3)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	jsoncliente4 := `{"Nome":"David","Email":"david@email.com","CPF":12345678900,"Pais":"EUA"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsoncliente4), &cliente4)
	fmt.Println(cliente4)
}