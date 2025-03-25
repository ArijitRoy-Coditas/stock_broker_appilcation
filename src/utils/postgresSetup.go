package utils

import (
	"fmt"
	"stock_broker_application/src/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresClient *models.DatabaseConfiguration

func InitPostgresConfg(configPath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("postgres")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Viper config paths:", viper.ConfigFileUsed())
		return fmt.Errorf("failed to read the config file: %s", err)
	}

	var postgresConfig models.PostgresConfig

	if err := viper.Unmarshal(&postgresConfig); err != nil {
		return fmt.Errorf("failed to unmarshal the config file %s", err)
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", postgresConfig.Host, postgresConfig.Port, postgresConfig.Username, postgresConfig.Password, postgresConfig.DBName, postgresConfig.SSLMode, postgresConfig.Timezone)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect wit the postgres: %s", err)
	}

	setDBInstance(db)
	return nil
}

func setDBInstance(db *gorm.DB) {
	postgresClient = &models.DatabaseConfiguration{GormDB: db}
}

func GetPostgresClient() *models.DatabaseConfiguration {
	return postgresClient
}
