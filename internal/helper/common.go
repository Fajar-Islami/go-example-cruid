package helper

import (
	"github.com/go-playground/validator/v10"
)

type ErrorStruct struct {
	Err  error
	Code int
}

var Validate = validator.New()
