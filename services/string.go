package services

import (
	"errors"
	"strings"
)

type StringServices interface {
	UpperCase(string) (string, error)
	Count(string) int
}

var errEmpty = errors.New("Empty String")

type StrService struct{}

func (s StrService) UpperCase(str string) (string, error) {
	if str == "" {
		return "", errEmpty
	}
	return strings.ToUpper(str), nil
}

func (s StrService) Count(str string) int {
	return len(str)
}
