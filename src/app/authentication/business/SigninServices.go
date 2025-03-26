package business

import (
	"authentication/commons/constants"
	"authentication/models"
	"authentication/repository"
	"context"
	"errors"
	"fmt"
	genericConstants "stock_broker_application/src/constants"
	genericModels "stock_broker_application/src/models"
	"stock_broker_application/src/utils"

	"github.com/gin-gonic/gin"
)

type SigninUserServices struct {
	signinUserRepository repository.SignInUserRepository
}

func NewSigninUserService(signinUserRepository repository.SignInUserRepository) *SigninUserServices {
	return &SigninUserServices{
		signinUserRepository: signinUserRepository,
	}
}

func (service *SigninUserServices) SignIn(ctx *gin.Context, spanCtx context.Context, bffSigninUserRequest models.BFFSinginUserRequest) (*genericModels.User, string, string, error) {
	postgresClient := utils.GetPostgresClient()
	tx := postgresClient.GormDB.Begin()
	if tx.Error != nil {
		return nil, "", "", fmt.Errorf(genericConstants.ErrBeginTx, tx.Error)
	}

	condition := map[string]interface{}{
		constants.Fieldemail: bffSigninUserRequest.Email,
	}

	columns := []string{
		genericConstants.FieldUsername,
		genericConstants.FieldEmail,
		genericConstants.FieldPassword,
	}

	user, err := service.signinUserRepository.SignIn(ctx, tx, condition, columns)
	if err != nil {
		tx.Rollback()
		return nil, "", "", fmt.Errorf("%w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, "", "", fmt.Errorf(genericConstants.ErrCommitTx, err)
	}

	if !utils.CompareHashPassword(user.Password, bffSigninUserRequest.Password) {
		return nil, "", "", errors.New(constants.ErrInvalidEmailorPassword)
	}

	accessToken, refreshToken, err := utils.GenerateToken(user.Username)
	if err != nil {
		return nil, "", "", fmt.Errorf(constants.ErrTokenGenerationFailed, err)
	}

	return user, accessToken, refreshToken, nil
}
