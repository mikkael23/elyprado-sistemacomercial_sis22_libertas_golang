package models

type Produto struct {
	ID       int    `json:"idproduto"`
	Descricao     string `json:"descricao"`
	PrecoCusto    float64 `json:"precocusto"`
	PrecoVenda    float64 `json:"precovenda"`
	SaldoEstoque float64 `json:"saldoestoque"`
	CodBarras int `json:"codbarras"`
	IdMarca int `json:"idmarca"`
	
}
