package main // Changed to 'package model' by the YAML

import (
	"encoding/json"
	"testing"
)

func TestGeneratedModel(t *testing.T) {
	// Sample data matching your CTO
	data := `{"shipmentId":"123","weight":10.5}`
	value := &Shipment{}

	if err := json.Unmarshal([]byte(data), value); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if value.ShipmentId == "" {
		t.Error("Data binding failed: ShipmentId is empty")
	}
}
