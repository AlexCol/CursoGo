package src

import (
	"fmt"
)

func UseGo() {
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Minha string %d", i)
		go showMessage(message) //forma de chamar metodo como uma nova thread
	}

	//time.Sleep(time.Duration(time.Hour.Seconds() * float64(1000)))
}

func showMessage(message string) {
	fmt.Println(message)
}
