package dtos

type ServerErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
