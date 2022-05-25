package main

import "fmt"

type Carro struct {
	Marca  string
	Modelo string
	Ano    int
}

func (c Carro) Print() {
	fmt.Println(c)
}

func (c Carro) Acelerar() {
	fmt.Println("acelerando...")
}

func (c Carro) GetModelo() string {
	return c.Modelo
}

func mainEstruturas() {
	c := Carro{
		Marca:  "Ferrari",
		Modelo: "812 GTS",
		Ano:    2020,
	}

	c.Print()
	c.Acelerar()
	fmt.Println(c.GetModelo())
}
