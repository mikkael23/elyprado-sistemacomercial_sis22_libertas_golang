package models

type Conta struct {
	ID_Pagar     int    `json:"idpagar"`
	data         string `json:"data"`
	valor        string `json:"valor"`
	vencimento   string `json:"vencimento"`
	pagamento    string `json:"pagamento"`
	valorpago    string `json:"valorpago"`
	idfornecedor string `json:"idfornecedor"`
}
