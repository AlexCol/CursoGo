package main

import (
	"module/src"
)

func main() {
	//src.ExePanic() //se descomentar, a aplicação 'morre', pois vai executar o Panic (estilo )
	src.ExeDefer()
}
