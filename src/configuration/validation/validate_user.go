package validation

import (
	"encoding/json"
	"errors"

	customError "github.com/MrBarreto/CRUD-OpenTelemetry/src/configuration/custom_error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	trans    ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		trans, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, trans)
	}
}

func ValidateUserError(
	validation_err error,
) *customError.CustomError {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return customError.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []customError.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := customError.Causes{
				Message: e.Translate(trans),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return customError.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return customError.NewBadRequestError("Error trying to convert fields")
	}
}
