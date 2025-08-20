package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!controle de fluxo com if
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	var mensagem string
	notaAluno := rand.Float32() * 10
	if notaAluno >= 7 {
		mensagem = fmt.Sprintf("Aprovado com nota %.2f", notaAluno)
	} else if notaAluno >= 5 {
		mensagem = fmt.Sprintf("Recuperação com nota %.2f", notaAluno)
	} else {
		mensagem = fmt.Sprintf("Reprovado com nota %.2f", notaAluno)
	}
	fmt.Println(mensagem)

	goEhLegal := true
	if goEhLegal {
		fmt.Println("Go é uma linguagem de programação legal!")
	}

	if true && false {
		fmt.Println("Isso não será impresso, pois a condição é falsa.")
	} else {
		fmt.Println("Isso será impresso, pois a condição do if é falsa.")
	}

	if true || false {
		fmt.Println("Isso será impresso, pois a condição do if é verdadeira.")
	} else {
		fmt.Println("Isso não será impresso, pois a condição do if é verdadeira.")
	}

	//erro, pois em go não tem if ternário
	//var ternario = true ? "Verdadeiro" : "Falso"

	//for i := 0; i < 5; i++ {
	for i := range 5 {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é ímpar")
		}
	}

	//! simulando forEach
	lista := []string{"Go", "Python", "Java"}
	for indice, item := range lista {
		fmt.Printf("Índice: %d, Valor: %s\n", indice, item)
	}

	//! simulando forEach
	idades := map[string]int{"Ana": 30, "Bob": 25}
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

	//! simulando While
	contador := 0
	for contador < 5 {
		fmt.Println("Contador:", contador)
		contador++
	}

	//! simulando Do While
	contador = 0
	for {
		fmt.Println("DoWhile:", contador)
		contador++
		if contador >= 5 {
			break
		}
	}
}
