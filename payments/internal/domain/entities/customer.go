package entities

type Customer struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	CardNumber   string         `json:"cardNumber"`
	Transactions []*Transaction `json:"transactions"`
}
