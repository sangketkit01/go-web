package main

import "testing"

var test = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expected-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range test {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error.")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error.")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %.2f but got %.2f ", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have.")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have.")
	}
}
