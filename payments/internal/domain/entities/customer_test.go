package entities

import (
	"testing"
)

func TestBalance(t *testing.T) {
	customerStr := Customer[string]{
		Id:         "1234567890",
		Name:       "John",
		CardNumber: "1234567890",
	}

	if customerStr.Name == "" {
		t.Error("can't create an Customer String object")
	}

	customerInt := Customer[int64]{
		Id:         1234567890,
		Name:       "John",
		CardNumber: "1234567890",
	}

	if customerInt.Name == "" {
		t.Error("can't create an Customer Int64 object")
	}
}
