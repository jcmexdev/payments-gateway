package entities

import (
	"testing"
)

func TestBalance(t *testing.T) {
	customerStr := Customer{
		Id:         "1234567890",
		Name:       "John",
		CardNumber: "1234567890",
	}

	if customerStr.Name == "" {
		t.Error("can't create an Customer String object")
	}
}
