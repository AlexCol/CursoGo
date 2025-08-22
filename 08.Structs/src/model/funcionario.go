package model

type Funcionario struct {
	Pessoa  //assim funciona a herança em go
	Cargo   string
	Salario float64
}
