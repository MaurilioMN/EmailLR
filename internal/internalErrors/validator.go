package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationStruct(obj interface{}) error {

	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validateErrors := err.(validator.ValidationErrors)
	validateError := validateErrors[0]

	field := strings.ToLower(validateError.StructField())

	switch validateError.Tag() {
	case "required":
		return errors.New(field + " is required " + validateError.Param())
	case "min":
		return errors.New(field + " is required with min " + validateError.Param())
	case "max":
		return errors.New(field + " is required with max " + validateError.Param())
	case "email":
		return errors.New(field + " is invalid! " + validateError.Param())
	}
	return nil
}
