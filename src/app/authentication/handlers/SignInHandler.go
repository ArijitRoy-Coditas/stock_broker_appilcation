package handlers

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"encoding/json"
	"errors"
	"net/http"

	genericConstants "stock_broker_application/src/constants"
	genericModels "stock_broker_application/src/models"
	"stock_broker_application/src/utils/validations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SigninUserHandler struct {
	service business.SigninUserServices
}

// HandleUserSignin handles user authentication and token generation.
// @Summary Sign in a user
// @Description Authenticates a user and returns access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.BFFSinginUserRequest true "User credentials"
// @Success 200 {object} models.SuccessAPIResponse
// @Failure 400 {object} models.ErrorAPIResponse
// @Failure 401 {object} models.ErrorAPIResponse
// @Failure 500 {object} models.ErrorAPIResponse
// @Router /api/auth/signin [post]
func NewSigninUserHandler(service business.SigninUserServices) *SigninUserHandler {
	return &SigninUserHandler{
		service: service,
	}
}

func (controller *SigninUserHandler) HandleUserSignin(ctx *gin.Context) {
	var bffSigninUserRequest models.BFFSinginUserRequest
	if err := ctx.ShouldBindJSON(&bffSigninUserRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	if err := validations.GetBFFValidator().Struct(&bffSigninUserRequest); err != nil {
		validationErros, _ := validations.FormatValidationErrors(err)
		ctx.IndentedJSON(http.StatusBadRequest, validationErros)
		return
	}

	user, accessToken, refreshToken, err := controller.service.SignIn(ctx, ctx.Request.Context(), bffSigninUserRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsgs := genericModels.
				ErrorMessage{
				Key:          constants.EmailorPasswordField,
				ErrorMessage: constants.ErrInvalidEmailorPassword,
			}

			response := genericModels.ErrorAPIResponse{
				Message: errorMsgs,
				Error:   constants.ErrInvalidPayload,
			}
			ctx.IndentedJSON(http.StatusUnauthorized, response)
			return
		}

		if err.Error() == constants.ErrInvalidEmailorPassword {
			errorMsgs := genericModels.ErrorMessage{
				Key:          constants.EmailorPasswordField,
				ErrorMessage: constants.ErrInvalidEmailorPassword,
			}

			response := genericModels.ErrorAPIResponse{
				Message: errorMsgs,
				Error:   constants.ErrInvalidPayload,
			}
			ctx.IndentedJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.IndentedJSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
			Message: genericModels.ErrorMessage{
				Key:          genericConstants.Interal,
				ErrorMessage: genericConstants.ErrInternalServer,
			},
			Error: constants.ErrAuthenticationFailed,
		})
	}

	ctx.SetCookie(
		constants.Name,
		refreshToken,
		constants.Time,
		constants.Path,
		constants.Domain,
		constants.Secure,
		constants.HttpOnly,
	)

	ctx.IndentedJSON(http.StatusOK, models.SuccessAPIResponse{
		Message: constants.UserLoggedInSuccessMsg,
		Data: models.BFFSigninUserResponse{
			Username:    user.Username,
			Email:       user.Email,
			AccessToken: accessToken,
		},
	})
}
