package _const

import (
	"errors"
	"fmt"
)

var InvalidValue = errors.New("Invalid Param ")

type CustomError struct {
	error      string
	fieldName  string
	fieldValue string
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("%s: field: %s - value: %s", ce.error, ce.fieldName, ce.fieldValue)
}
