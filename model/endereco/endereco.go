package model

type Endereco struct {
	Rua         string `db:"rua" json:"rua"`
	Cep         string `db:"cep" json:"cep"`
	Numero      int    `db:"numero" json:"numero"`
	Bairro      string `db:"bairro" json:"bairro"`
	Complemento string `db:"complemento" json:"complemento"`
}
