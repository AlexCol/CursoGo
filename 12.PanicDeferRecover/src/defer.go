package src

import (
	"fmt"
	"os"
)

func ExeDefer() {
	defer showMessaWithDefer() //!
	fmt.Println("Iniciando exemplo")
	file, error := os.Create("./settings.txt")

	if error != nil {
		panic(error) //se der problema para criar um arquivo, 'explode' erro que se não for tratado, vai parar a aplicação (mas o defer ainda é chamado)
	}
	defer file.Close() //!

	file.Write([]byte("Teste"))
	//file.Close() //pode ser chamado no fim do arquivo
	fmt.Println("Finalizando exemplo")
}

func showMessaWithDefer() {
	fmt.Println("Chamado com Defer") //como o defer foi anunciado antes do panic, esse será executado mesmo que panic ocorra
}

//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// Tdo defer registrado até o ponto de execução será chamado no momento em que a função terminar,
// seja por retorno normal ou por panic. Se ocorrer um panic, os defer são executados em ordem
// inversa ao registro (LIFO), garantindo que recursos abertos até aquele ponto sejam liberados.
// Porém, se o panic acontecer antes do defer ser declarado, ele não será registrado e portanto
// não será executado.
//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
