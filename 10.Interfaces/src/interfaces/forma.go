package interfaces

type Forma interface {
	Area() float64
	//Teste()
}

// interface não precisa ser 'herdada', ela funciona automaticamente,
// desde que a classe concreta implemente TODOS
// metodos da interface
//* ex: aqui temos Forma, e na pasta Model temos Circulo e Retangulo, ambas
//* implementam Area, portanto podem ser chamadas no lugar de qualquer
//* metodo que chame Forma (usado em main)
//? se descomentar Teste, já apresenta erro, pois circulo e retangulo
//? não implementam Teste
