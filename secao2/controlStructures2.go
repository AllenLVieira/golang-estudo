package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	marcador := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			marcador = false
			break
		}
		fmt.Println(i)
	}
	fmt.Println(marcador)

	dia := "Sexta"
	switch dia {
	case "Sexta":
		fmt.Println("Hoje é sexta")
	case "Segunta", "Terça", "Quarta":
		fmt.Println("Dia chato")
	default:
		fmt.Println("Padrão")
	}

	switch {
	case dia == "Sexta":
		fmt.Println("Hoje é sexta")
		break
	}
}
