package admin_controller_util

import (
	"fmt"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
)

func ValidateChallengeAndTop(challenges []admin_request.Challenge, tops []admin_request.Top) *rest_err.RestErr {
	var causes []rest_err.Causes

	for i, ch := range challenges {
		prefix := fmt.Sprintf("challenges[%d]", i)

		if ch.Challenge == "" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".challenge",
				Message: "Challenge é obrigatório",
			})
		}
		if ch.Category == "" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".category",
				Message: "Categoria é obrigatória",
			})
		} else if ch.Category != "iniciante" && ch.Category != "scaled" && ch.Category != "rx" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".category",
				Message: "Categoria inválida",
			})
		}
		if ch.Sex == "" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".sex",
				Message: "Sexo é obrigatório",
			})
		} else if ch.Sex != "M" && ch.Sex != "F" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".sex",
				Message: "Sexo deve ser M ou F",
			})
		}
	}

	for i, top := range tops {
		prefix := fmt.Sprintf("tops[%d]", i)

		if top.Top == 0 {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".top",
				Message: "Top é obrigatório",
			})
		}
		if top.Category == "" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".category",
				Message: "Categoria é obrigatória",
			})
		} else if top.Category != "iniciante" && top.Category != "scaled" && top.Category != "rx" {
			causes = append(causes, rest_err.Causes{
				Field:   prefix + ".category",
				Message: "Categoria inválida",
			})
		}
	}

	if len(causes) > 0 {
		return &rest_err.RestErr{
			Message: "Validation Error",
			Err:     "validation_error",
			Code:    400,
			Causes:  causes,
		}
	}

	return nil
}