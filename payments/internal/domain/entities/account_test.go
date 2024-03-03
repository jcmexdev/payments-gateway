package entities

import (
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	transaction := Transaction{
		Id:                 "1234567890",
		Amount:             100.0,
		OriginAccount:      "1234567890",
		DestinationAccount: "1234567890",
		CreatedAt:          time.Now(),
	}

	if transaction.Id == "" {
		t.Error("can't create a Transaction object")
	}
}
