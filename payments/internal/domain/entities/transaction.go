package entities

import "time"

type Transaction struct {
	Id                 string    `json:"id" bson:"_id,omitempty"`
	Amount             float64   `json:"amount" bson:"amount"`
	OriginAccount      string    `json:"originAccount" bson:"originAccount"`
	DestinationAccount string    `json:"destinationAccount" bson:"destinationAccount"`
	Type               string    `json:"type" bson:"type"`
	CreatedAt          time.Time `json:"createdAt" bson:"createdAt"`
}
