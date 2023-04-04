package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateVariable() bool {
	validate = validator.New()
	email := "jfjdlfjle@tencent.com"
	errs := validate.Var(email, "required,email")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}
