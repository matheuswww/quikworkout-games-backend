package admin_request

import (
	"regexp"
	"time"

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

	errors = append(errors, Validator.RegisterTranslation("category", *translator, func(ut ut.Translator) error {
		return ut.Add("category", "categoria deve ser iniciante, scaled ou rx", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("category", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterTranslation("closing_date", *translator, func(ut ut.Translator) error {
		return ut.Add("closing_date", "data encerramento inválida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("closing_date", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterTranslation("start_date", *translator, func(ut ut.Translator) error {
		return ut.Add("start_date", "data de inicio inválida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("start_date", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterTranslation("time", *translator, func(ut ut.Translator) error {
		return ut.Add("time", "time inválido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("time", fe.Field())
		errors = append(errors, err)
		return t
	}))

	errors = append(errors, Validator.RegisterValidation("category", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if field == "" {
			return true
		}
		if field != "iniciante" && field != "scaled" && field != "rx" {
			return false
		}
		return true
	}))

	errors = append(errors, Validator.RegisterValidation("closing_date", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		dateRegex := `^20\d\d-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`
		matched, _ := regexp.MatchString(dateRegex, dateStr)
		if !matched {
			return false
		}
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return false
		}
		parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local)
		
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return parsedDate.After(today)
	}))

	errors = append(errors, Validator.RegisterValidation("start_date", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		dateRegex := `^20\d\d-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`
		matched, _ := regexp.MatchString(dateRegex, dateStr)
		if !matched {
			return false
		}
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return false
		}
		parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local)

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

		return !parsedDate.Before(today)
	}))


	errors = append(errors, Validator.RegisterValidation("time", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if(val == "") {
			return true
		}
		pattern := `^([01]\d|2[0-3]):[0-5]\d:[0-5]\d\.\d{3}$`
		matched, _ := regexp.MatchString(pattern, val)
		return matched
	}))
}