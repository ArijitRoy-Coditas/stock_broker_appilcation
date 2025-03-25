package validations

import (
	"fmt"
	"regexp"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

func ValidatePasswordConstraints(password string) []string {
	var errors []string

	if len(password) < 8 {
		errors = append(errors, constants.ErrPasswordMinLength)
	}
	if !regexp.MustCompile(constants.LowercaseRegex).MatchString(password) {
		errors = append(errors, constants.ErrPasswordLowercase)
	}
	if !regexp.MustCompile(constants.UppercaseRegex).MatchString(password) {
		errors = append(errors, constants.ErrPasswordUppercase)
	}
	if !regexp.MustCompile(constants.DigitRegex).MatchString(password) {
		errors = append(errors, constants.ErrPasswordDigit)
	}
	if !regexp.MustCompile(constants.SpecialCharRegex).MatchString(password) {
		errors = append(errors, constants.ErrPasswordSpecialChar)
	}

	return errors
}

func FormatValidationErrors(err error) ([]models.ErrorMessage, string) {
	var validationErrors []models.ErrorMessage
	var validationErrorsStr string

	for _, err := range err.(validator.ValidationErrors) {
		var errorMsg string
		fieldName := err.Field()
		// Handle required fields first
		if err.Tag() == "required" {
			fieldName = strings.ToLower(fieldName)
			errorMsg = fmt.Sprintf(constants.ErrFieldRequired, fieldName)
		} else {
			// Custom handling for specific fields
			switch err.Field() {
			case constants.FieldPassword:
				passwordErrors := ValidatePasswordConstraints(err.Value().(string))
				if len(passwordErrors) > 0 {
					for _, msg := range passwordErrors {
						validationErrors = append(validationErrors, models.ErrorMessage{
							Key:          err.Field(),
							ErrorMessage: msg,
						})
					}
					continue // Skip adding the generic message
				}
			case constants.FieldConfirmPassword:
				if err.Tag() == "eqfield" {
					errorMsg = constants.ErrConfirmPasswordMatch
				}
			case constants.FieldPanCard:
				errorMsg = constants.ErrInvalidPanCard
			case constants.FieldPhoneNumber:
				errorMsg = constants.ErrInvalidPhoneNumber
			default:
				errorMsg = fmt.Sprintf(constants.ErrInvalidValue, err.Field())
			}
		}

		validationErrors = append(validationErrors, models.ErrorMessage{
			Key:          fieldName,
			ErrorMessage: errorMsg,
		})
		validationErrorsStr += fieldName + " is invalid; "
	}

	return validationErrors, validationErrorsStr
}

func panCardValidator(f1 validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(constants.PANCardRegex, f1.Field().String())
	return matched
}

func strongPasswordValidator(f1 validator.FieldLevel) bool {
	re := regexp2.MustCompile(constants.PasswordRegex, 0) // Compile regex with PCRE support
	matched, _ := re.MatchString(f1.Field().String())
	return matched
}

func init() {
	bffValidator = validator.New()
	bffValidator.RegisterValidation("panCard", panCardValidator)
	bffValidator.RegisterValidation("strongPassword", strongPasswordValidator)
}

func GetBFFValidator() *validator.Validate {
	return bffValidator
}
