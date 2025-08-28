package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4}
	reverse(mySlice)
	fmt.Println(mySlice)

	mySlice2 := []string{"A", "B", "C", "D"}
	reverse(mySlice2)
	fmt.Println(mySlice2)
}

// func reverse[T int | string](slice []T) {
// func reverse[T any](slice []T) {
func reverse[T myContraint](slice []T) {
	i := 0
	for i < len(slice)/2 {
		esquerda := i
		direita := len(slice) - 1 - i
		slice[esquerda], slice[direita] = slice[direita], slice[esquerda]
		i++
	}
}

type myContraint interface {
	int | string
}
