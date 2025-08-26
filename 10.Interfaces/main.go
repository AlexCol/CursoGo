package main

import (
	"fmt"
	"module/src/interfaces"
	"module/src/model"
)

func main() {
	circulo := model.Circulo{
		Raio: 5,
	}
	MostraArea(circulo)

	retangulo := model.Retangulo{
		Largura: 10,
		Altura:  10,
	}
	MostraArea(retangulo)
}

func MostraArea(f interfaces.Forma) {
	area := f.Area()
	fmt.Printf("A área é: %.2f\n", area)
}
