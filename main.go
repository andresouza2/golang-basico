package main

import (
	"errors"
	"fmt"
	"log"
)


func main() {
	
	// aqui o erro é ignorado e não tratado aqui
	// res, _ := soma(5, 5)

	res, err := soma(5, 5)

	if err != nil {
		log.Fatal((err.Error()))
	}

	fmt.Println(res)

}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
} 