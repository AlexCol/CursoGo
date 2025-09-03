package src

import (
	"fmt"
	"time"
)

func UseMutex() {
	//var wg sync.WaitGroup

	i := 0
	for x := 0; x < 100; x++ {
		go ChangeNumber(&i, x)
	}

	//wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Printf("Valor i é: %d", i)
}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}
