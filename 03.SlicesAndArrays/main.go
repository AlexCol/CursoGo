package main

import "fmt"

func main() {
	//!! slice ('arrays' dinâmicos - declarados sem tamanho fixo)
	lista := []int{1, 2, 3, 4, 5} //
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

	//!!array (array de tamanho fixo)
	fmt.Println("array------------------------")
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	array[0] = 10
	fmt.Println(array)
	//array = append(array, 6) //não funciona, pois o array é fixo, não pode crescer, pra isso precisa usar slice

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

	meuOutroMap := map[int]string{1: "Um", 2: "Dois", 3: "Três", 4: "Quatro"} //A ordem de armazenamento e iteração será sempre imprevisível.
	for chave, valor := range meuOutroMap {
		fmt.Println(chave, valor)
	}
	valor, exists := meuOutroMap[5] //busca por chave inexistente
	if !exists {
		fmt.Println("Chave não encontrada")
	} else {
		fmt.Println("Valor:", valor)
	}

	mapForLoop := make(map[int]string)
	itensMap := 5
	for i := 0; i < itensMap; i++ {
		mapForLoop[i] = fmt.Sprintf("Número %d", i)
	}
	mapForLoop[itensMap+1] = fmt.Sprintf("Número %d", itensMap+1)

	i := 0
	for i < itensMap+2 {
		valor, existe := mapForLoop[i]
		if existe {
			fmt.Printf("Chave %d encontrada com valor: %s\n", i, valor)
		} else {
			fmt.Printf("Chave %d não encontrada\n", i)
		}
		i++
	}

	//!range
	listaForRange := []int{1, 2, 3, 4, 5, 6, 7}
	for index, valor := range listaForRange {
		fmt.Printf("Index: %d, Valor: %d\n", index, valor)
	}

	//!delete (so funciona para map)
	meuMapForDelete := map[int]string{1: "Um", 2: "Dois", 3: "Três", 4: "Quatro"}
	fmt.Println(meuMapForDelete)
	delete(meuMapForDelete, 2)
	fmt.Println("Após delete:")
	fmt.Println(meuMapForDelete)

	//!set (é um map, onde o item é a chave, o valor do item é vazio ou qualquer coisa, pois é ignorado)
	meuSet := make(map[string]struct{})
	meuSet["item1"] = struct{}{}
	meuSet["item2"] = struct{}{}
	meuSet["item3"] = struct{}{}
	fmt.Println(meuSet)

}

/*
forEach:
	for 'indice', 'valor' := range 'colecao' {
		...logica
	}

for (normal)
for i := 0; i < len(listaForRange); i++ {
	...logica
}

while
for 'minhaCondicao' {
	...logica
}

Do While
for { //sem condição
	...logica
	if ('minhaCondicao') { //if simula o 'while' no fim do código
		break
	}
}
*/
