package util

import (
	"github.com/go-playground/validator/v10"
)

type FieldErrorEntry struct {
	Name  string `json:"name"`
	Err   string `json:"error"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
}

type ValidationErr[T any] struct {
	Message string                     `json:"message"`
	Fields  map[string]FieldErrorEntry `json:"fields"`
	Data    *T                         `json:"data"`
}

var val = validator.New()

func GetFieldErr[T any](err *ValidationErr[T], field string) string {
	if err == nil {
		return ""
	}

	fieldErr, ok := err.Fields[field]

	if !ok {
		return ""
	}

	return fieldErr.Tag
}

func GetErrData[T any](err *ValidationErr[T]) *T {
	if err == nil || err.Data == nil {
		return new(T)
	}
	return err.Data
}

func Validate[T any](value *T) (*ValidationErr[T], error) {
	err := val.Struct(value)
	errs, ok := err.(validator.ValidationErrors)

	if ok {
		fields := make(map[string]FieldErrorEntry)

		for _, field := range errs {
			fields[field.Field()] = FieldErrorEntry{
				Name:  field.Field(),
				Err:   field.Error(),
				Tag:   field.Tag(),
				Value: field.Value(),
			}
		}

		result := &ValidationErr[T]{
			Message: err.Error(),
			Fields:  fields,
			Data:    value,
		}

		return result, nil
	}

	if err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return nil, err
	}

	return nil, nil
}

func (err *ValidationErr[T]) Error() string {
	return err.Message
}
