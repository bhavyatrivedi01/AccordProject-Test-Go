package main // The YAML's 'sed' command will change this to the correct package name

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAllGeneratedModels(t *testing.T) {
	testCases := []struct {
		Name  string
		Value interface{}
		Data  string
	}{
		{
			Name:  "Shipment", 
			Value: &Shipment{}, 
			Data:  `{"shipmentId":"123","weight":10.5}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			val := reflect.ValueOf(tc.Value)
			// Reflection helps ensure the generated code matches expected Go types
			if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
				t.Fatalf("%s is not a pointer to a struct", tc.Name)
			}

			if err := json.Unmarshal([]byte(tc.Data), tc.Value); err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tc.Name, err)
			}

			_, err := json.Marshal(tc.Value)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tc.Name, err)
			}
		})
	}
}
