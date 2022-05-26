package main

import "fmt"

func Algo(algo interface{}) {
	fmt.Println(algo)
}

func mainInterfaces2() {
	Algo(2.44)
	Algo("Allen")
	Algo(struct{}{})

	mapa := make(map[string]interface{})
	mapa["Nome"] = "Allen"
	mapa["Idade"] = 25

	fmt.Println(mapa)
}
