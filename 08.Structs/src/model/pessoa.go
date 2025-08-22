package model

import "time"

type Pessoa struct {
	Nome      string
	sobrenome string // campo privado, não acessível fora do pacote
	Endereco  Endereco
	DataNasc  time.Time
}

func (p *Pessoa) SetSobrenome(sobrenome string) { //precisa ser ponteiro em (p *Pessoa) para que as alterações sejam refletidas na struct original
	p.sobrenome = sobrenome
}

func (p Pessoa) NomeCompleto() string { //(p Pessoa) é uma 'copia de pessoa'
	return p.Nome + " " + p.sobrenome
}

func (pessoa Pessoa) CalculaIdade() int {
	anoAtual := time.Now().Year()
	anoNasc := pessoa.DataNasc.Year()
	return anoAtual - anoNasc
}
