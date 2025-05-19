package get_custom_validator

import (
	"errors"
	"log"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translation "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	Validator  = validator.New()
	Translator ut.Translator
)

func CustomValidator(obj any) (ut.Translator, error) {
	err := Validator.Struct(obj)
	return Translator, err
}

func init() {
	pt := pt_BR.New()
	unt := ut.New(pt, pt)
	var found bool
	Translator, found = unt.GetTranslator("pt_BR")
	if !found {
		logger.Error("Translator not found", errors.New("translator not found"), zap.String("journey", "Init translator"))
		log.Fatal(errors.New("translator not found"))
	}
	pt_translation.RegisterDefaultTranslations(Validator, Translator)
}
