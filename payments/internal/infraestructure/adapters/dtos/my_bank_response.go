package dtos

type MyBankResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}
