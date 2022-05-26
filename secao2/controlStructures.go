package main

import "fmt"

func mainControlStructures() {
	x := true
	y := &x

	if y == nil {
		fmt.Println("Value is null")
	} else if *y {
		fmt.Println("Value is true")
	} else {
		fmt.Println("Value is false")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []string{"allen", "de lima", "vieira"}

	for i, valor := range arr {
		fmt.Println(i)
		fmt.Println(valor)
	}

	mapa := make(map[string]interface{})
	mapa["Nome"] = "Allen"
	mapa["Idade"] = 25

	for k, v := range mapa {
		fmt.Printf("Chave: %s, Valor: %v\n", k, v)
	}
}
