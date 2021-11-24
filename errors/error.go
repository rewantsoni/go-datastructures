package errors

import "fmt"

var IndexOutOfBoundError = func(i int) error {
	return fmt.Errorf("index %d is out of bound", i)
}

var NotFondError = func(e interface{}, list interface{}) error {
	return fmt.Errorf("element %v not found in the %v", e, list)
}

var InvalidOperationError = func(reason string) error {
	return fmt.Errorf("invalid operation: %s", reason)
}

var EmptyListError = InvalidOperationError("list is empty")

var TypeMismatchError = func(expected, got string) error {
	return fmt.Errorf("expected %s got %s", expected, got)
}