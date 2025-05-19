package default_validator

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translation "github.com/go-playground/validator/v10/translations/pt"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt := pt_BR.New()
		unt := ut.New(pt, pt)
		var found bool
		translator, found = unt.GetTranslator("pt_BR")
		if !found {
			logger.Error("Translator not found", errors.New("translator not found"), zap.String("journey", "Init translator"))
			log.Fatal(errors.New("translator not found"))
		}
		pt_translation.RegisterDefaultTranslations(val, translator)
	}
}

func HandleDefaultValidatorErrors(validation_err error) *rest_err.RestErr {
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
