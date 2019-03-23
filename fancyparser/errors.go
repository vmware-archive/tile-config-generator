package fancyparser

import "fmt"

type InvalidIndexTypeError struct {
	ProvidedType IndexType
	RequiredType IndexType
}

func (e InvalidIndexTypeError) Error() string {
	return fmt.Sprintf("Expected type: %s, but got: %s", e.RequiredType, e.ProvidedType)
}

type NoValueAtEndOfIndexError struct {
	RemainingValue interface{}
}

func (e NoValueAtEndOfIndexError) Error() string {
	return fmt.Sprintf("%v left to traverse", e.RemainingValue)
}
