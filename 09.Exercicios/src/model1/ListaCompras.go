package model1

import (
	"errors"
	"time"
)

type ListaCompras struct {
	Mercado    string
	DataCompra time.Time
	Itens      []ItensCompra
}

func NewList(mercado string) (*ListaCompras, error) {
	if mercado == "" {
		return nil, errors.New("nome do mercado não pode ser vazio")
	}

	novaLista := ListaCompras{
		Mercado:    mercado,
		DataCompra: time.Now(),
		Itens:      []ItensCompra{},
	}
	return &novaLista, nil
}

func (lc *ListaCompras) AddItem(item ItensCompra) {
	lc.Itens = append(lc.Itens, item)
}
