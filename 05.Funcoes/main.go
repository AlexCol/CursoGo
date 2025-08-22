package main

import "fmt"

/*
A função init() em Go é executada automaticamente antes da função main(), assim que o pacote é inicializado.
Ou seja, quando você roda o programa, init() é chamada primeiro, depois main().
Você não precisa (nem pode) chamar init() manualmente.
Se houver múltiplos arquivos no mesmo pacote, a ordem entre arquivos não é garantida (depende do compilador).
*/
func init() {
	fmt.Println("Inicializando o main...")
}

func main() {
	n1 := 1
	n2 := 2
	soma := soma(n1, n2) //! exemplo soma
	fmt.Printf("O valor da soma é: %d\n", soma)
	fmt.Printf("O valor de n1 é: %d\n", n1)

	//!exemplo de função de múltiplos valores
	soma, subtracao, multiplicacao, divisao := operacoes(n1, n2)
	fmt.Printf("O valor da soma é: %d\n", soma)
	fmt.Printf("O valor da subtração é: %d\n", subtracao)
	fmt.Printf("O valor da multiplicação é: %d\n", multiplicacao)
	if n2 != 0 {
		fmt.Printf("O valor da divisão é: %f\n", divisao)
	} else {
		fmt.Println("Divisão por zero não é permitida.")
	}

	//!exemplo de função que retorna uma função
	fmt.Println("Soma de ", n1, "e", n2, "é:", operacoes2("soma")(n1, n2))
	fmt.Println("Subtração de ", n1, "e", n2, "é:", operacoes2("subtracao")(n1, n2))
	fmt.Println("Multiplicação de ", n1, "e", n2, "é:", operacoes2("multiplicacao")(n1, n2))
	fmt.Println("Divisão de ", n1, "e", n2, "é:", operacoes2("divisao")(n1, n2))

	//!exemplo de uma função que executa outra função
	minhaFuncao := func() {
		fmt.Println("Esta é uma função de callback.")
	}
	CallBack(minhaFuncao)

}

//! exemplo simples de função
func soma(a int, b int) int {
	n1 := 210
	fmt.Printf("O valor de n1 dentro de soma é: %d\n", n1)
	return a + b
}

//! exemplo de função que retorna vários valores
//func operacoes(n1 int, n2 int) (int, int, int, float64) {
func operacoes(n1 int, n2 int) (soma int, subtracao int, multiplicacao int, divisao float64) { //! nomeando os parâmetros de retorno (facilita para quem for consumir a função)
	soma = n1 + n2
	subtracao = n1 - n2
	multiplicacao = n1 * n2
	if n2 != 0 {
		divisao = float64(n1) / float64(n2)
	}
	return soma, subtracao, multiplicacao, divisao
}

//! exemplo de função que retorna uma função
func operacoes2(operacao string) func(int, int) float64 {
	switch operacao {
	case "soma":
		return func(a int, b int) float64 {
			return float64(a + b)
		}
	case "subtracao":
		return func(a int, b int) float64 {
			return float64(a - b)
		}
	case "multiplicacao":
		return func(a int, b int) float64 {
			return float64(a * b)
		}
	case "divisao":
		return func(a int, b int) float64 {
			if b != 0 {
				return float64(a) / float64(b)
			}
			return 0
		}
	default:
		return nil
	}
}

//! exemplo de uma função que executa outra função
func CallBack(callback func()) {
	callback()
}
