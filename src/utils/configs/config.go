package configs

import (
	"fmt"
	"os"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"

	"github.com/spf13/viper"
)

func LoadConfig(configPath, configFile, configType string) (*models.PostgresConfig, error) {
	var postgresConfig models.PostgresConfig

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)
	fmt.Println(os.Getwd())
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf(constants.ErrReadConfigFailed, err)
	}

	if err := viper.Unmarshal(&postgresConfig); err != nil {
		return nil, fmt.Errorf(constants.ErrUnmarshallConfigFailed, err)
	}

	return &postgresConfig, nil
}
