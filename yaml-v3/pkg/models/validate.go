package models

import (
	"errors"
	"fmt"
	"strings"
)

type Constraint func() error

type Validator struct {
	constraints []Constraint
}

func (v *Validator) Validate() error {
	error_list := []error{}

	for _, constraint := range v.constraints {
		err := constraint()
		if err != nil {
			error_list = append(error_list, err)
		}
	}

	return errors.Join(error_list...)
}

func (v *Validator) AddConstraint(c Constraint) {
	v.constraints = append(v.constraints, c)
}

// Check if a string is empty
func NotEmpty(value string, fieldName string) Constraint {
	return func() error {
		if len(strings.TrimSpace(value)) == 0 {
			return fmt.Errorf("%s cannot be empty", fieldName)
		}
		return nil
	}
}

// NotZeroLength checks if a slice length is zero
func NotZeroLength[T any](value []T, fieldName string) Constraint {
	return func() error {
		if len(value) == 0 {
			return fmt.Errorf("%s cannot be zero length", fieldName)
		}
		return nil
	}
}

func PositiveNotZero[T int | int64](value T, fieldName string) Constraint {
	return func() error {
		if value <= 0 {
			return fmt.Errorf("%s cannot be equal to or less then zero", fieldName)
		}
		return nil
	}
}
