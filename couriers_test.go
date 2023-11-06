package tracking51

import (
	"context"
	"testing"
)

func TestGetAllCouriers(t *testing.T) {
	key := "you api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	response, err := client.GetAllCouriers(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*[]Courier)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}
}
