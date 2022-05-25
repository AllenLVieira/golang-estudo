package main

import (
	"fmt"
	"strings"
)

func main() {

	builder := strings.Builder{}
	var texto string
	texto = "Meu nome é"
	fmt.Println(texto)
	texto2 := " Allen de Lima Vieira"
	fmt.Println(texto2)
	builder.WriteString(texto)
	builder.WriteString(texto2)
	fmt.Println(builder.String())
}

/* func main() {
	var inteiro int
	inteiro = 10
	fmt.Println(inteiro)

	//Inicialização
	var (
		x = 1
		y = 2
	)
	fmt.Println(x + y)

	// Cast de variaveis, são inicializadas automaticamente como 0
	var m int32
	var n int64
	fmt.Println(int64(m) + n)
} */
