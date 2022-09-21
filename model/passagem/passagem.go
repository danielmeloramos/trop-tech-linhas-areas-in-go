package model

type Passagem struct {
	Origem      string  `db:"origem" json:"origem"`
	Destino     string  `db:"destino" json:"destino"`
	DataOrigem  string  `db:"dataOrigem" json:"dataOrigem"`
	DataDestino string  `db:"dataDestino" json:"dataDestino"`
	Valor       float64 `db:"valor" json:"valor"`
}
