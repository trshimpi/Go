package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-error", 100.0, 0.0, 0.0, true},
	{"expect-5", 100.0, 20.0, 5.0, false},
	{"expect-fraction", 100.0, 3.0, 33.333333, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Expected error, got nil")
			}
		} else {
			if err != nil {
				t.Error("Expected no error, got ", err)
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f, got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestDivide2(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
