package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kialkuz/task-manager/pkg/logger"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ErrorFormatter struct {
	messages map[string]func(fe validator.FieldError) string
}

func NewErrorFormatter() *ErrorFormatter {
	return &ErrorFormatter{
		messages: map[string]func(fe validator.FieldError) string{
			"required": func(fe validator.FieldError) string {
				return "is required"
			},
			"email": func(fe validator.FieldError) string {
				return "must be a valid email"
			},
			"min": func(fe validator.FieldError) string {
				return fmt.Sprintf("must be at least %s characters", fe.Param())
			},
			"taskdate": func(fe validator.FieldError) string {
				return "date format 20060102"
			},
		},
	}
}

func (f *ErrorFormatter) ViewFormat(err error) []FieldError {
	var out []FieldError

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range f.processValidationErrors(errs) {
			out = append(out, FieldError{
				Field:   e["field"].(string),
				Message: e["message"].(string),
			})
		}
	}

	return out
}

func (f *ErrorFormatter) PrepareForLogs(err error) logger.Field {
	out := make(logger.Field)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range f.processValidationErrors(errs) {
			out[e["field"].(string)] = e["message"].(string)

		}
	}

	return out
}

func (f *ErrorFormatter) processValidationErrors(errs validator.ValidationErrors) []map[string]any {
	var out []map[string]any

	for _, e := range errs {
		msgFunc, exists := f.messages[e.Tag()]

		var message string
		if exists {
			message = msgFunc(e)
		} else {
			message = "invalid value"
		}

		out = append(out, map[string]any{
			"field":   e.Field(),
			"message": message,
		})
	}

	return out
}
