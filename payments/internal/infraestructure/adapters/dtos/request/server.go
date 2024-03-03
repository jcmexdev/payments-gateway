package request

type TransferRequest struct {
	OriginAccount      string  `json:"originAccount"`
	DestinationAccount string  `json:"destinationAccount"`
	Amount             float64 `json:"amount"`
}
