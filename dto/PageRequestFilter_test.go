package dto

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFilterValue_Export(t *testing.T) {
	tests := []struct {
		name     string
		rawJson  string
		expected interface{}
	}{
		{
			name:     "String value",
			rawJson:  `"2023-10-01"`,
			expected: "2023-10-01",
		},
		{
			name:     "Integer value",
			rawJson:  `125`,
			expected: 125,
		},
		{
			name:     "Decimal value",
			rawJson:  `125.5`,
			expected: 125.5,
		},
		{
			name:     "Boolean true",
			rawJson:  `true`,
			expected: true,
		},
		{
			name:     "String array",
			rawJson:  `["active", "pending"]`,
			expected: []string{"active", "pending"},
		},
		{
			name:     "Integer array",
			rawJson:  `[1, 2, 3]`,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Float array",
			rawJson:  `[1.1, 2.2, 3.3]`,
			expected: []float64{1.1, 2.2, 3.3},
		},
		{
			name:     "Null value",
			rawJson:  `null`,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fv FilterValue
			_ = json.Unmarshal([]byte(tt.rawJson), &fv)

			got := fv.Export()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Export() type %T = %v, want type %T = %v", got, got, tt.expected, tt.expected)
			}
		})
	}
}

func TestFilterValue_ExportAs(t *testing.T) {
	tests := []struct {
		name    string
		rawJson string
		testFn  func(fv FilterValue)
	}{
		{
			name:    "ExportAs String",
			rawJson: `"hello"`,
			testFn: func(fv FilterValue) {
				val := ExportAs[string](fv)
				if val != "hello" {
					t.Errorf("Expected hello, got %v", val)
				}
			},
		},
		{
			name:    "ExportAs Int",
			rawJson: `42`,
			testFn: func(fv FilterValue) {
				val := ExportAs[int](fv)
				if val != 42 {
					t.Errorf("Expected 42, got %v", val)
				}
			},
		},
		{
			name:    "ExportAs Float64",
			rawJson: `42.5`,
			testFn: func(fv FilterValue) {
				val := ExportAs[float64](fv)
				if val != 42.5 {
					t.Errorf("Expected 42.5, got %v", val)
				}
			},
		},
		{
			name:    "ExportAs String Slice",
			rawJson: `["a", "b"]`,
			testFn: func(fv FilterValue) {
				val := ExportAs[[]string](fv)
				expected := []string{"a", "b"}
				if !reflect.DeepEqual(val, expected) {
					t.Errorf("Expected %v, got %v", expected, val)
				}
			},
		},
		{
			name:    "ExportAs Int Slice",
			rawJson: `[1, 2, 3]`,
			testFn: func(fv FilterValue) {
				val := ExportAs[[]int](fv)
				expected := []int{1, 2, 3}
				if !reflect.DeepEqual(val, expected) {
					t.Errorf("Expected %v, got %v", expected, val)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fv FilterValue
			err := json.Unmarshal([]byte(tt.rawJson), &fv)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}
			tt.testFn(fv)
		})
	}
}
