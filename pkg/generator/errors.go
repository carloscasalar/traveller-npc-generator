package generator

import "fmt"

type InvalidListError struct {
	listName string
}

func newInvalidListError(listName string) InvalidListError {
	return InvalidListError{listName: listName}
}

func (e InvalidListError) Error() string {
	return fmt.Sprintf("invalid list %s, a non-empty list must be provided", e.listName)
}
