package terminal

const packageName = "terminal"

// Error holds the operation and the cause. `cause` must not be nil.
type Error struct {
	operation string
	cause     error
}

// Error returns the package, the operation, and the cause of the error.
func (e *Error) Error() string {
	return packageName + ": " + e.operation + ": " + e.cause.Error()
}

// Unwrap returns the cause of the error.
func (e *Error) Unwrap() error {
	return e.cause
}
