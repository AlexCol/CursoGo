package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var ponteiro *int

	ponteiro = &a

	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de ponteiro:", ponteiro)
	fmt.Println("Valor apontado por ponteiro:", *ponteiro)

	*ponteiro = 30
	fmt.Println("Novo valor de a:", a)

	ponteiro = &b
	*ponteiro = 40
	fmt.Println("Novo valor de b:", b)

	c := *ponteiro
	fmt.Println("Valor de c:", c)

	mudaValor(ponteiro) //ponteiro aponta pra o local de memoria de b
	fmt.Println("Novo valor de b:", b)
	fmt.Println("Valor de c:", c)
}

func mudaValor(ponteiro *int) {
	*ponteiro = 26
}
