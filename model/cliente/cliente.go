package model

type Cliente struct {
	Cpf          int64  `db:"cpf" json:"cpf"`
	PrimeiroNome string `db:"primeiroNome" json:"primeiroNome"`
	SobreNome    string `db:"sobreNome" json:"sobreNome"`
	NomeCompleto string `db:"nomeCompleto" json:"nomeCompleto"`
	EnderecoId   int    `db:"enderecoId" json:"enderecoId"`
}
