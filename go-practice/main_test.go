package main

import "testing"

func TestDivide(t *testing.T) {
	got, err := divide(6, 3)
	want := 2
	if got != want || err != nil {
		t.Errorf("divide(6, 3) = %d, want %d, error '%v'", got, want, err)
	}
}
