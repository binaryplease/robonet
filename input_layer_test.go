package robonet

import "testing"

func TestInputLayer_Calculate(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lay := &InputLayer{
				LayerFields: tt.fields.LayerFields,
			}
			lay.Calculate()
		})
	}
}
