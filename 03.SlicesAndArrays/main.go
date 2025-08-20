package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println(lista)
	lista = append(lista, 6) //mas
	fmt.Println(lista)

	subLista := make([]int, 0, 10)
	for _, valor := range lista {
		if valor > 3 {
			subLista = append(subLista, valor)
		}
	}
	fmt.Println(subLista)

	//!! slice
	//Primeiro índice inclusivo, segundo exclusivo.
	fmt.Println("slice------------------------")
	sliced := lista[1:4] // [2 3 4] //de index 1 até index 4-1 (é exclusivo)
	fmt.Println("sliced:", sliced)
	sliced = lista[:3] // [1 2 3] //até index 3-1 (é exclusivo)
	fmt.Println("sliced:", sliced)
	sliced = lista[2:] // [3 4 5 6] //a partir de index 2
	fmt.Println("sliced:", sliced)

	//!! referencia
	fmt.Println("referencia------------------------")
	s := make([]int, 2, 4) // len=2, cap=4
	s2 := s

	s = append(s, 10) // len=3, cap=4 -> mesmo array subjacente, não aloca novo
	fmt.Println(s)    // [0 0 10]
	fmt.Println(s2)   // [0 0] -> s2 ainda tem len=2, mas aponta pro mesmo array

	s2 = s2[:len(s)] // 'forçando' s2 a ver o tamanho adequado
	fmt.Println(s2)  // [0 0 10] -> agora s2 "enxerga" o terceiro elemento

	s = append(s, 20, 30) // len=5, cap=4 -> capacidade insuficiente, cria novo array
	fmt.Println(s)        // [0 0 10 20 30]
	fmt.Println(s2)       // [0 0 10] -> s2 continua apontando para o array antigo

	// s2 = s2[:len(s)] // daria panic: len(s2)=3, cap(s2)=4, não pode ir além da cap

	//!!map
	fmt.Println("map------------------------")
	type pessoa struct {
		nome  string
		idade int
	}

	mapPessoas := make(map[string]pessoa)
	var nomePessoa string
	var newPessoa pessoa

	newPessoa = pessoa{"Ana", 28}
	nomePessoa = newPessoa.nome
	mapPessoas[nomePessoa] = newPessoa

	newPessoa = pessoa{"João", 30}
	nomePessoa = newPessoa.nome
	mapPessoas[nomePessoa] = newPessoa

	newPessoa = pessoa{"Maria", 25}
	nomePessoa = newPessoa.nome
	mapPessoas[nomePessoa] = newPessoa

	fmt.Println(mapPessoas)
	for chave, valor := range mapPessoas {
		fmt.Println(chave, valor)
	}
}
