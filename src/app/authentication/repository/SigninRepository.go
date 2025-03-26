package repository

import (
	"authentication/commons/constants"
	"context"
	genericModels "stock_broker_application/src/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type signInUserRepository struct{}

type SignInUserRepository interface {
	SignIn(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, columns []string) (*genericModels.User, error)
}

func NewSignInUser() *signInUserRepository {
	return &signInUserRepository{}
}

func (user *signInUserRepository) SignIn(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, columns []string) (*genericModels.User, error) {
	start := time.Now()
	logger := logrus.New()

	var User genericModels.User

	result := db.WithContext(ctx).Select(columns).Debug().Where(conditions).First(&User)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	logger.WithFields(logrus.Fields{
		"result":  result,
		"latency": time.Since(start).Milliseconds(),
	}).Info(constants.UserLoggedInSuccessMsg)

	return &User, nil
}
