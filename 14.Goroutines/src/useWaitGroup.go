package src

import (
	"fmt"
	"sync"
	"time"
)

func UseWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3) // informa que 3 goroutines precisam finalizar antes do Wait ser liberado

	go callApi(&wg)         // inicia uma nova goroutine para chamar a API
	go callDatabase(&wg)    // inicia uma nova goroutine para chamar o Database
	go internalProcess(&wg) // inicia uma nova goroutine para o processo interno

	wg.Wait() // bloqueia até que todas as goroutines chamem Done() (contador chega a zero)

	fmt.Println("Todos os processo finalizados!")
}

func callDatabase(wg *sync.WaitGroup) {
	fmt.Println("Chamando Database")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando chamada a Database")
	wg.Done() // sinaliza que esta goroutine terminou (decrementa o contador em 1)
}

func callApi(wg *sync.WaitGroup) {
	fmt.Println("Chamando API")
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizando chamada a API")
	wg.Done() // sinaliza que esta goroutine terminou (decrementa o contador em 1)
}

func internalProcess(wg *sync.WaitGroup) {
	fmt.Println("Chamando InternalProcess")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando chamada a InteernalProcess")
	wg.Done() // sinaliza que esta goroutine terminou (decrementa o contador em 1)
}
