package main

import (
	"fmt"
	"log"
	"net/http"
)

// import (
// 	"curso-go/math"
// 	"fmt"
// )

func main() {
	// a := 34
	// b := "andre"
	// c := 3.14
	// d := true
	// e := `teste \n linha 3`

	// fmt.Printf("%v %T\n", a, a)
	// fmt.Printf("%v %T\n", b, b)
	// fmt.Printf("%v %T\n", c, c)
	// fmt.Printf("%v %T\n", d, d)
	// fmt.Printf("%v %T\n", e, e)

	// resultado := math.Soma(2, 4)
	// resultado2 := math.SomaX(45)

	// fmt.Printf("%v\n", resultado)
	// fmt.Printf("%v\n", resultado2)

	res, err := http.Get("http://googlesssscom.br")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%v\n", res.Header)
}
