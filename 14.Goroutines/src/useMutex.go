package src

import (
	"fmt"
	"sync"
	"time"
)

func UseMutex() {
	//var wg sync.WaitGroup

	var m sync.Mutex
	i := 0
	for x := 0; x < 100000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}

	//wg.Wait()

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Valor i é: %d", i)
}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}
