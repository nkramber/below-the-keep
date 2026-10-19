package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		nums        [2]int
		want        int
		expectError bool
	}{
		{[2]int{4, 2}, 2, false},
		{[2]int{10, 2}, 5, false},
		{[2]int{5, 0}, 0, true},
	}

	for _, testCase := range testCases {
		a, b := testCase.nums[0], testCase.nums[1]
		got, err := divide(a, b)
		want := testCase.want
		if got != want {
			t.Errorf("divide(%d, %d) = %d, want %d", a, b, got, want)
			continue
		}
		if testCase.expectError && err == nil {
			t.Errorf("divide(%d, %d) = %d, wanted error, got %v", a, b, got, err)
		}
		if err != nil && !testCase.expectError {
			t.Errorf("divide(%d, %d) = %d, did not expect an error, got '%v'", a, b, got, err)
		}
	}
}

func TestDivideCause(t *testing.T) {
	got, err := divide(6, 0)
	want := 0

	if err == nil {
		t.Fatal("divide(6, 0): wanted an error, got nil")
	}
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("divide(6, 0): error %v does not hold %v", err, ErrDivideByZero)
	}
	if got != want {
		t.Errorf("divide(6, 0): got %d, want %d", got, want)
	}
}
