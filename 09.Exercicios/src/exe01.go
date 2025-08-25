package exercicios

import (
	"fmt"
	"modulo/src/model1"
)

func Ex01() {
	novaLista, error := model1.NewList("Mercado X")
	if error != nil {
		fmt.Println(error)
		return
	}

	item1 := model1.ItensCompra{DescricaoItem: "CocaCola", Preco: 10.5}
	novaLista.AddItem(item1)

	item2 := model1.ItensCompra{DescricaoItem: "Macarrão", Preco: 4.75}
	novaLista.AddItem(item2)

	fmt.Printf("Compra realizada em %s.\n", novaLista.Mercado)
	fmt.Printf("No dia %s.\n", novaLista.DataCompra.Format("02/01/2006 15:04:05"))
	fmt.Println("Itens da lista:")
	for _, item := range novaLista.Itens {
		fmt.Printf("%s por R$%.2f\n", item.DescricaoItem, item.Preco)
	}
}
