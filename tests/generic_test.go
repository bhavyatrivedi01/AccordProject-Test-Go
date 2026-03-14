package main

import (
    "encoding/json"
    "reflect"
    "testing"
)

// This test uses reflection to find and test all structs in the package
func TestAllGeneratedModels(t *testing.T) {
    // List of "testable" concepts you expect in your generated code
    // In a real GSoC project, you could automate this list via 'concerto parse'
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
            // Use reflection to verify the type
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
