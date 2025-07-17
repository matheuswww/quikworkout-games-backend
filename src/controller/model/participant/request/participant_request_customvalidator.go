package participant_request

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

	errors = append(errors, Validator.RegisterTranslation("sex", *translator, func(ut ut.Translator) error {
		return ut.Add("sex", "sex must be M or F", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("sex", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterValidation("sex", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if val == "" {
			return true
		}
		sexRegex := regexp.MustCompile(`^[MF]$`)
		return sexRegex.MatchString(val)
	}))
}