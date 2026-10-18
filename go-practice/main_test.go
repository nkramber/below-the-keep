package main

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		nums      [2]int
		want      int
		errorText string
	}{
		{[2]int{4, 2}, 2, ""},
		{[2]int{10, 2}, 5, ""},
		{[2]int{5, 0}, 0, "cannot divide by zero"},
	}

	for _, testCase := range testCases {
		a, b := testCase.nums[0], testCase.nums[1]
		got, err := divide(a, b)
		want := testCase.want
		if got != want || (err != nil && testCase.errorText == "") {
			t.Errorf("divide(%d, %d) = %d, want %d, error '%v'", a, b, got, want, err)
			continue
		}
		if testCase.errorText != "" && err == nil {
			t.Errorf("divide(%d, %d) = %d, wanted error, got %v", a, b, got, err)
		}
		if err != nil && err.Error() != testCase.errorText {
			t.Errorf("divide(%d, %d) = %d, wanted error text '%v', got '%v'", a, b, got, testCase.errorText, err.Error())
		}
	}
}
