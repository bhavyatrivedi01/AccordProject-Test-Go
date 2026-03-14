package main // Changed by workflow

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGeneratedModel(t *testing.T) {
	// Sample data matching the CTO definition
	data := `{"shipmentId":"123","weight":10.5}`
	
	// We use the struct name directly.
	// The workflow moves this file INTO the folder where Shipment is defined.
	value := &Shipment{}

	t.Run("Unmarshal", func(t *testing.T) {
		if err := json.Unmarshal([]byte(data), value); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}
	})

	t.Run("ReflectCheck", func(t *testing.T) {
		v := reflect.ValueOf(value).Elem()
		idField := v.FieldByName("ShipmentId")
		if !idField.IsValid() || idField.String() != "123" {
			t.Error("Reflection failed to find or verify ShipmentId")
		}
	})
}
