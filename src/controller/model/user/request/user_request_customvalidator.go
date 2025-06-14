package user_request

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
)

var (
	Validator  = validator.New()
	translator ut.Translator
)

func CustomValidator(obj any) (*ut.Translator, error) {
	err := Validator.Struct(obj)
	return &translator, err
}

func init() {
	translator := &get_custom_validator.Translator
	Validator := get_custom_validator.Validator
	var errors []error

	errors = append(errors, Validator.RegisterTranslation("cpf", *translator, func(ut ut.Translator) error {
		return ut.Add("cpf", "CPF inválido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("cpf", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterValidation("cpf", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) == 11 {
			return ValidateCpf(fl.Field().String())
		}
		return false
	}))

	errors = append(errors, Validator.RegisterTranslation("user", *translator, func(ut ut.Translator) error {
		return ut.Add("user", "usuário inválido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("user", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterValidation("user", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()
		if len(username) == 0 || len(username) > 30 {
			return false
		}
		if strings.HasPrefix(username, ".") {
			return false
		}
		if strings.HasSuffix(username, ".") {
			return false
		}
		if strings.Contains(username, "..") {
			return false
		}
		for _, ch := range username {
			if (ch >= 'a' && ch <= 'z') ||
				(ch >= 'A' && ch <= 'Z') ||
				(ch == 'ç' || ch == 'Ç') ||
				(ch >= '0' && ch <= '9') ||
				ch == '.' || ch == '_' {
				continue
			}
			return false
		}
		return true
	}))
}
