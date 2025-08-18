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
}
