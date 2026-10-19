package terminal

import (
	"errors"
	"testing"
)

const testOperation string = "testOperation"
const testCause string = "test error for cause"

func TestError(t *testing.T) {
	err := &Error{operation: testOperation, cause: errors.New(testCause)}

	want := "terminal: testOperation: test error for cause"
	got := err.Error()

	if want != got {
		t.Errorf("&Error{%v, %v} = '%v', want '%v'", testOperation, testCause, got, want)
	}
}

func TestUnwrap(t *testing.T) {
	causeErr := errors.New(testCause)
	err := &Error{operation: testOperation, cause: causeErr}

	if !errors.Is(err, causeErr) {
		t.Errorf("errors.Is(%q, %q) = false, want true; errors.Unwrap(err) = %v", err, causeErr, errors.Unwrap(err))
	}
}
