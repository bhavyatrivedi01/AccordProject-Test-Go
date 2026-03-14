// We change this to a generic package name because the script 
// will inject it into the generated folders.
package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestAllGeneratedModels now uses reflection to test whatever 
// structs are available in the current package scope.
func TestAllGeneratedModels(t *testing.T) {
	testCases := []struct {
		Name  string
		Value interface{}
		Data  string
	}{
		{
			// Note: This must match a struct name in your .cto file
			Name:  "Shipment", 
			Value: &Shipment{}, 
			Data:  `{"shipmentId":"123","weight":10.5}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			val := reflect.ValueOf(tc.Value)
			if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
				t.Fatalf("%s is not a pointer to a struct", tc.Name)
			}

			// Test Unmarshal
			if err := json.Unmarshal([]byte(tc.Data), tc.Value); err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tc.Name, err)
			}

			// Test Marshal
			_, err := json.Marshal(tc.Value)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tc.Name, err)
			}
		})
	}
}
