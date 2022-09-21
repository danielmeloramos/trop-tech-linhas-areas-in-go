package model

type Viagem struct {
	CodigoReserva   string  `db:"codigoReserva" json:"codigoReserva"`
	IdaVolta        bool    `db:"idaVolta" json:"idaVolta"`
	DataCompra      string  `db:"dataCompra" json:"dataCompra"`
	ValorTotal      float64 `db:"valorTotal" json:"valorTotal"`
	ResumoViagem    string  `db:"resumoViagem" json:"resumoViagem"`
	PassagemIdaId   int     `db:"passagemIdaId" json:"passagemIdaId"`
	PassagemVoltaId int     `db:"passagemVoltaId" json:"passagemVoltaId"`
	ClienteId       int     `db:"clienteId" json:"clienteId"`
}
