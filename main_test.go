package main

import "testing"

func Test_tested(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"tested"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tested()
		})
	}
}
