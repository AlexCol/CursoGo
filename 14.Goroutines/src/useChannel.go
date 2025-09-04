package src

import "fmt"

func UseChannel() {
	//channel := make(chan int)
	channel := make(chan int, 50) //ao informar um numero, ele vai mandar para o recebedor, apenas quando tiver esse numero de itens
	go setList(channel)

	for valor := range channel {
		fmt.Println("recebendo:", valor)
	}
}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("enviando:", i)
		channel <- i
	}
	close(channel)
}

//!
//se ao declarar 'chan' fizar isso:
// ? '<-chan' significa que ele é somente leitura
// ? 'chan<-' signigica que ele é somente escrita
//? sem nada ele é escrita e leitura
