package models


type Domicilio struct {
	Calle string `json:"calle"`
	Numero string	`json:"numero"`
	Ciudad string `json:"ciudad"`
	Referencia string `json:"referencia"`
}