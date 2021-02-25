package main

import (
	"fmt"
)

type Veiculo interface{
	buzinar()
	getNome() string
	getAno() int
}

type Pessoa struct{
	Nome string
	Veiculo Veiculo
}

type Carro struct{
	Nome string
	Ano int
}

func (c Carro) buzinar(){
	fmt.Println(c.Nome, " buzinou!")
}

func (c Carro) getNome() string {
	return c.Nome
}

func (c Carro) getAno() int {
	return c.Ano
}

type Moto struct{
	Nome string
	Ano int
}

func (m Moto) buzinou(){
	fmt.Println(m.Nome, " buzinou!")
}

func (m Moto) getNome() string {
	return m.Nome
}

func (m Moto) getAno() int {
	return m.Ano
}


func main() {
	carroTop := Carro{
		Nome: "Renegade",
		Ano: 2019}

	john := Pessoa{Nome: "Johntas", Veiculo: carroTop}
	fmt.Println(john.Nome, "tem um", john.Veiculo.getNome(), "ano", john.Veiculo.getAno())
}