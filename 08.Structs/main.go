package main

import (
	"fmt"
	aliasTipo "modulo/src/model"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	endereco := aliasTipo.Endereco{
		Rua:    "Rua A",
		Cidade: "Cidade B",
	}
	fmt.Println("Endereço:", endereco)

	pessoa := aliasTipo.Pessoa{
		Nome:     "Alexandre",
		Endereco: endereco,
		DataNasc: time.Date(1985, 6, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println("Endereço:", pessoa)

	idade := pessoa.CalculaIdade()
	fmt.Println("Idade:", idade)

	pessoa.SetSobrenome("Coletti")
	fmt.Println("Nome completo:", pessoa.NomeCompleto())

	funcionario := aliasTipo.Funcionario{
		Pessoa:  pessoa,
		Cargo:   "Desenvolvedor",
		Salario: 5000.00,
	}
	funcionario.SetSobrenome("Trabalho")                    //meu sobrenome é trabalho kkkkk
	fmt.Println("Funcionário:", funcionario.NomeCompleto()) //continua funcionando, pois funcionario herda de pessoa

	//criando funcionário completo
	funcionario2 := aliasTipo.Funcionario{
		Pessoa: aliasTipo.Pessoa{
			Nome: "Maria",
			Endereco: aliasTipo.Endereco{
				Rua:    "Rua B",
				Cidade: "Cidade A",
			},
			DataNasc: time.Date(1990, 5, 15, 0, 0, 0, 0, time.Local),
		},
		Cargo:   "Analista",
		Salario: 6000.00,
	}
	fmt.Println(funcionario2)
	funcionario2.SetSobrenome("Silva")
	fmt.Println("Funcionário 2:", funcionario2.NomeCompleto())

}
