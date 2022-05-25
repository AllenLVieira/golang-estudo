package main

import "fmt"

func main() {
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
}
