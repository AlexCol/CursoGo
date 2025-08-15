package codes

import (
	"fmt"
	"modulo/src/types"
)

const ExportVar = 5

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func ReadInput() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)
}

func Soma() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error reading first number:", err)
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error reading second number:", err)
		return
	}

	result := soma(num1, num2)
	fmt.Println("The sum is:", result)
}

func LerDoisNumeros() {
	var num1, num2 int
	n, err := fmt.Scanln(&num1, &num2)
	if err != nil {
		fmt.Println("Error reading numbers:", err)
		return
	}

	fmt.Println("Quantidade de numeros lidos:", n)
	fmt.Println("Valores lidos:", num1, "-", num2)

}

func CriaPessoa() {
	endereco := types.Endereco{
		Rua:    "Rua Exemplo",
		Numero: 123,
		Cidade: "Cidade Exemplo",
		Estado: "Estado Exemplo",
	}

	pessoa := types.Pessoa{
		Nome:     "Nome Exemplo",
		Idade:    30,
		Endereco: endereco,
	}

	fmt.Println("Pessoa:", pessoa)
}

// private
func soma(n1, n2 int) int {
	return n1 + n2
}
