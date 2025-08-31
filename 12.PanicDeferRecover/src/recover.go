package src

import (
	"fmt"
	"os"
)

func ExeRecover() {
	//defer func() {}() //formato para criar uma função anonima, e já invocar ela
	//defer myRecover()

	readFile()

	fmt.Println("Fim")
}

func readFile() {
	// geralmente é bom colocar o recover na função onde o erro pode ocorrer,
	// se deixar no metodo ExeRecover, por mais que ocorra o recover, o metodo
	// 'morre' e não continua  e fim não é executado
	defer myRecover()

	_, err := os.Open(".settings.txtt")
	if err != nil {
		panic("FileNotExists")
	}
}

func myRecover() {
	//if r := recover(); r != nil { forma opcional de if no go, ele permite inicializar uma variavel antes de realizar a comparação, abaixo fiz separado
	r := recover()
	if r != nil {
		fmt.Println("Erro capturado:", r)
	}
}
