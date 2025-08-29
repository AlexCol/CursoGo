package src

import (
	"fmt"
	"os"
)

func ExePanic() {
	/*
		fmt.Println("oi")
		panic("Processo que fecha o sistema")
		fmt.Println("tchau")
	*/

	//_, err := os.Open("c:/alexandre/gerar_mapa.py")
	_, err := os.Open("c:/alexandre/gerar_mapa.pyss")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo aberto com sucesso")
}
