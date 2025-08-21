package exercicios

import "fmt"

func Ex02() {
	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	somaMenores := 0
	somaMaiores := 0

	for _, item := range lista {
		if item > 5 {
			somaMaiores += item
		} else {
			somaMenores += item
		}
	}

	fmt.Printf("Soma dos menores ou iguais a 5: %d\n", somaMenores)
	fmt.Printf("Soma dos maiores que 5: %d\n", somaMaiores)
}
