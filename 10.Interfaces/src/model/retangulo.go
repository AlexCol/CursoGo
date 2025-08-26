package model

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (ret Retangulo) Area() float64 {
	return ret.Altura * ret.Largura
}
