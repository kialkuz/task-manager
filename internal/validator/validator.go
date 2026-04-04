package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
	datePkg "github.com/kialkuz/task-manager/pkg/date"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterValidation("taskdate", taskDateValidator)

	return &Validator{validate: v}
}

func (v *Validator) ValidateStructDTO(s any) error {
	return v.validate.Struct(s)
}

func taskDateValidator(fieldLevel validator.FieldLevel) bool {
	date := fieldLevel.Field().String()
	if date == "" {
		return true
	}

	_, err := time.Parse(datePkg.DateFormat, date)
	return err == nil
}
