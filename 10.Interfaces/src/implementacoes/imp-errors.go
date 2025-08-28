package implementacoes

import (
	"errors"
	"fmt"
)

func MyErrors() {
	var err error                                  //lint:ignore S1021 'exemplo de ignore de warning'
	err = errors.New(("Usando interface de erro")) //esse foi ignorado globalmente pelo arquivo "staticcheck.conf"
	//fmt.Println(err)
	exibeError(err)

	minhaNetError := networkError{net: false, hardware: true}
	exibeError(minhaNetError)
}

func exibeError(err error) { //metodo que espera algo que implemente error (que tenha o metodo Error())
	fmt.Println(err.Error())
}

// ? struct exemplo
type networkError struct {
	net      bool
	hardware bool
}

func (n networkError) Error() string {
	if n.net {
		return "Erro de rede"
	}
	if n.hardware {
		return "Problema de hardware"
	}
	return "Erro desconhecido"
}
