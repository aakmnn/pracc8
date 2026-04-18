package main

import "testing"

func TestSubtractTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 5, 3, 2},
		{"positive minus zero", 5, 0, 5},
		{"negative minus positive", -3, 2, -5},
		{"both negative", -5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := Divide(10, 2)
		if err != nil || res != 5 {
			t.Errorf("fail")
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Errorf("expected error")
		}
	})
}
