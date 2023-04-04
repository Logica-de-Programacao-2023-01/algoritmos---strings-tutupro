package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e
//converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

func main() {
	var x string // criação da string

	fmt.Print("Digite uma palavra com letras minusculas: ")
	fmt.Scan(&x) // Escaneamento da Palavra

	o := strings.ToUpper(x) // converção de strings minusculas para maiusculas

	fmt.Printf("A palavra que você digitou %s foi convertida para maiusculo, dando: %s", x, o)
}
