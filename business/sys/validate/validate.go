// Package validate contains the support for validating models.
package validate

import (
	"encoding/json"
	"errors"
)

var ErrInvalidID = errors.New("ID is not its proper from")

type ErrorResponse struct {
	Error  string `json:"error"`
	Fields string `json:"fields,omitempty"`
}

type RequestError struct {
	Err    error
	Status int
	Fields error
}

func NewRequestError(err error, status int) error {
	return &RequestError{Err: err, Status: status}
}

func (err *RequestError) Error() string {
	return err.Err.Error()
}

type FieldsError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type FieldsErrors []FieldsError

func (fe FieldsErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}
