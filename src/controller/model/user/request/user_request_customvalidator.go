package user_request

import (
	"regexp"

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

	errors = append(errors, Validator.RegisterTranslation("date", *translator, func(ut ut.Translator) error {
		return ut.Add("date", "data de nascimento inválida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("date", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		dateRegex := `^20\d\d-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01])$`
		return regexp.MustCompile(dateRegex).MatchString(fl.Field().String())
	}))

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
}
