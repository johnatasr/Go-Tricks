package main

import (
	"fmt"
)

func main() {
	
	nome := "Nome top"
	sb := []byte(nome)
	
	fmt.Printf("%v - %T - formated to bytes:  %v - %T \n", nome , nome, sb, sb)
	
	// Pega por inteiro 
	for _, v := range nome {
		// sbLetter := []byte(v)
		fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v )
	}
	
	fmt.Println("===================")
	
	// Pega por byte
	for i := 0; i < len(nome); i++ {
		fmt.Printf("%v - %T - %#U - %#x \n", nome[i], nome[i], nome[i], nome[i])
	}

}
