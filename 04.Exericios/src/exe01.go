package exercicios

func Ex01() {
	lista := []int{1, 4, 7, 8, 7, 3}
	soma := 0
	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}
	println("Soma:", soma)
}
