package model

type Products struct {
	Storage  int `json:"storage_id"`
	Product  int `json:"product_id"`
	Amount   int `json:"amount"`
	Reserved int `json:"reserved"`
}
