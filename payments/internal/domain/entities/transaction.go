package entities

import "time"

type Transaction struct {
	Id                 string
	Amount             float64
	OriginAccount      string
	DestinationAccount string
	Type               string
	CreatedAt          time.Time
}
