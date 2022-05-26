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
}
