package util

import (
	"github.com/fatih/structs"
	"github.com/go-playground/validator/v10"
)

type FieldErrorEntry struct {
	Name  string `json:"name"`
	Err   string `json:"error"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
}

type ValidationErr struct {
	Message string                     `json:"message"`
	Fields  map[string]FieldErrorEntry `json:"fields"`
	Data    map[string]any             `json:"data"`
}

var val = validator.New()

func GetFieldErr(err *ValidationErr, field string) string {
	if err == nil {
		return ""
	}

	fieldErr, ok := err.Fields[field]

	if !ok {
		return ""
	}

	return fieldErr.Tag
}

func GetErrData(err *ValidationErr) map[string]any {
	if err == nil || err.Data == nil {
		return make(map[string]any)
	}
	return err.Data
}

func GetErrDataStr(err *ValidationErr, field string) string {
	if err == nil || err.Data == nil {
		return ""
	}
	val, ok := err.Data[field]
	if !ok {
		return ""
	}
	return val.(string)
}

func Validate[T any](value *T) (*ValidationErr, error) {
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

		result := &ValidationErr{
			Message: err.Error(),
			Fields:  fields,
			Data:    structs.Map(value),
		}

		return result, nil
	}

	if err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return nil, err
	}

	return nil, nil
}

func (err *ValidationErr) Error() string {
	return err.Message
}
