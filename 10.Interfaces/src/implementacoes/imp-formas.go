// Package implementacoes para implementacoes e separacao dos exemplos
package implementacoes

import (
	"fmt"
	"module/src/interfaces"
	"module/src/model"
)

func Formas() {
	circulo := model.Circulo{
		Raio: 5,
	}
	mostraArea(circulo)

	retangulo := model.Retangulo{
		Largura: 10,
		Altura:  10,
	}
	mostraArea(retangulo)
}

// ? exemplo de area
func mostraArea(f interfaces.Forma) {
	area := f.Area()
	fmt.Printf("A área é: %.2f\n", area)
}
