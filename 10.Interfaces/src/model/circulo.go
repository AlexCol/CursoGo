package model

import "math"

type Circulo struct {
	Raio float64
}

func (cir Circulo) Area() float64 {
	return math.Pi * math.Pow(cir.Raio, 2)
}
