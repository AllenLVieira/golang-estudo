package main

import "fmt"

type Carros interface {
	Acelerar()
	Frear()
}

type Ferrari struct {
	ModeloFerrari string
}

type Maserati struct {
	ModeloMaserati string
}

func NovaFerrari(arg string) Carros {
	return &Ferrari{arg}
}

func (f *Ferrari) Acelerar() {
	fmt.Println("Ferrari está acelerando")
	fmt.Println(f.ModeloFerrari)
}

func (f *Ferrari) Frear() {
	fmt.Println("Ferrari está freando")
	fmt.Println(f.ModeloFerrari)
}

func (m *Maserati) Acelerar() {
	fmt.Println("Maserati está acelerando")
	fmt.Println(m.ModeloMaserati)
}

func mainInterfaces() {
	f := NovaFerrari("812 GTS")
	m := Maserati{"Quattroporte"}
	f.Acelerar()
	m.Acelerar()
}
