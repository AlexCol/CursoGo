package src

import (
	"fmt"
	"time"
)

func UseWaitGroup() {
	go callApi()
	go callDatabase()
	go internalProcess()
}

func callDatabase() {
	fmt.Println("Chamando Database")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando chamada a Database")
}

func callApi() {
	fmt.Println("Chamando API")
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizando chamada a API")
}

func internalProcess() {
	fmt.Println("Chamando InternalProcess")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando chamada a InteernalProcess")
}
