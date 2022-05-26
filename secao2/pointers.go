package main

import (
	"fmt"
)

func trocar(m, n *int) {
	var temp int
	temp = *n
	*n = *m
	*m = temp
}

func mainPointers() {
	x := 5
	ptr := &x
	fmt.Printf("Endereço na memória: ")
	fmt.Println(ptr)
	fmt.Printf("Valor: ")
	fmt.Println(*ptr)

	m, n := 5, 10
	fmt.Println(m, n)
	trocar(&m, &n)
	fmt.Println(m, n)
}
