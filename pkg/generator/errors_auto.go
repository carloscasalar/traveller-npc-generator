// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package generator

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// EmptyListErr returns the error.
func EmptyListErr(listName string) error {
	return fmt.Errorf(`[ERR-1] invalid %v list: should contain at least one item`, listName)
}

// EmptyListErrWrap wraps the error.
func EmptyListErrWrap(
	listName string,
	err error,
) error {
	return errors.Wrap(err, "[ERR-1] invalid %v list: should contain at least one item")
}

// GeneratorErrorsType represents the error type.
type GeneratorErrorsType int

const (

	// EmptyListErrType represents the error type for EmptyListErr.
	EmptyListErrType GeneratorErrorsType = iota
	// GeneratorErrorsUnknownType represents unknown type for GeneratorErrors
	GeneratorErrorsUnknownType
)

// ListGeneratorErrors returns the list of errors.
func ListGeneratorErrors() []string {
	return []string{"[ERR-1] invalid %v list: should contain at least one item"}
}

// IdentifyGeneratorErrors checks the identity of an error
func IdentifyGeneratorErrors(err error) GeneratorErrorsType {
	errStr := err.Error()
	switch {
	case strings.HasPrefix(errStr, "[ERR-1]"):
		return EmptyListErrType
	default:
		return GeneratorErrorsUnknownType
	}
}
