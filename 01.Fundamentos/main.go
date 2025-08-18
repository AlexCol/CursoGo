package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func main() {
	//codes.HelloWorld()
	//codes.ReadInput()
	//codes.Soma()
	//codes.CriaPessoa()
	//codes.LerDoisNumeros()
	//fmt.Println("Valor da variável exportada:", codes.ExportVar)

	//tipos de dados
	var inteiro int = 10
	fmt.Println("Valor inteiro:", inteiro)

	var pontoFlutuante float64 = 3.14
	fmt.Println("Valor ponto flutuante:", pontoFlutuante)

	var booleano bool = true
	fmt.Println("Valor booleano:", booleano)

	var texto string = "Olá, Mundo!"
	fmt.Println("Valor texto:", texto)

	var data time.Time = time.Now()
	fmt.Println("Valor data:", data)

	var myArr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Valor array:", myArr)

	//operações
	var n1, n2 float64
	fmt.Println("Aplicando operações matemáticas...")
	n1, n2 = 30.5, 20.3

	// Adição
	soma := n1 + n2
	fmt.Println("Soma:", soma)

	// Subtração
	subtracao := n1 - n2
	fmt.Println("Subtração:", subtracao)

	// Multiplicação
	multiplicacao := n1 * n2
	fmt.Println("Multiplicação:", multiplicacao)

	// Divisão
	if n2 != 0 {
		divisao := n1 / n2
		fmt.Println("Divisão:", divisao)
	} else {
		fmt.Println("Divisão: Não é possível dividir por zero.")
	}

	// Módulo
	modulo := 10 % 2
	fmt.Println("Módulo:", modulo)

	// Gerar números aleatórios
	for i := 1; i <= 5; i++ {
		numeroAleatorio := rand.Intn(100)
		fmt.Println("Número aleatório:", numeroAleatorio)
	}

	//ver o tipo de uma variavel (se int, string, bool)
	//!forma 1
	varTipada := 5
	var tipo string = fmt.Sprintf("%T", varTipada)
	fmt.Println("Tipo da variável:", tipo)

	//!forma 2
	varTipada2 := "Olá"
	tipo2 := reflect.TypeOf(varTipada2)
	fmt.Println("Tipo da variável:", tipo2)

	//uma constante
	const pi = 3.14
	fmt.Println("Valor de pi:", pi)
	//pi = 3

	//casting
	var int8Var int8 = 127
	var float64Var float64 = float64(int8Var)
	fmt.Println("Valor int8:", int8Var)
	fmt.Println("Valor float64:", float64Var)

	float64Var += 1
	int8Var = int8(float64Var)          //vai dar overflow
	fmt.Println("Valor int8:", int8Var) //por isso aqui vai retornar -128

	float64Var += 1.5
	int32Var := int32(float64Var) // vai truncar, pois int32 não tem casas decimais
	fmt.Println("Valor int32:", int32Var)
	fmt.Println("Valor float64:", float64Var)

	//casting from string
	numFromString, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Erro ao converter string para int:", err)
	} else {
		fmt.Println("Valor from string:", numFromString)
	}

	boolFromString, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Erro ao converter string para bool:", err)
	} else {
		fmt.Println("Valor from string:", boolFromString)
	}
	//fmt.Scanln()
}
