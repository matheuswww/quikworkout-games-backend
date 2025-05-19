package custom_validator

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
)

func HandleCustomValidatorErrors(translator ut.Translator, validation_err error) *rest_err.RestErr {
	var jsonValidationError validator.ValidationErrors
	if errors.As(validation_err, &jsonValidationError) {
		var errorsCauses []rest_err.Causes
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("alguns campos são invalidos", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("erro tentando converter os campos")
	}
}
