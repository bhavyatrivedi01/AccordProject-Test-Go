package org_example_shipping

import (
        "encoding/json"
        "testing"
)

func TestGeneratedModel(t *testing.T) {
        data := `{
                "shipmentId": "123",
                "weight": 10.5,
                "status": "PLANNED"
        }`

        var value Shipment
        if err := json.Unmarshal([]byte(data), &value); err != nil {
                t.Fatalf("failed to unmarshal: %v", err)
        }

        if value.ShipmentId != "123" {
                t.Fatalf("expected ShipmentId to be 123, got %q", value.ShipmentId)
        }

        if value.Weight != 10.5 {
                t.Fatalf("expected Weight to be 10.5, got %v", value.Weight)
        }
}
