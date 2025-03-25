package main

import (
	ServiceConstants "authentication/commons/constants"
	"authentication/router"
	"fmt"
	"log"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	err := utils.InitPostgresConfg("../../config")
	if err != nil {
		log.Fatalf(constants.ErrDBConnectionFailed, err)
	}
	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info(fmt.Sprintf(constants.RunningServerPort, ServiceConstants.PortDefaultValude))
	router.Run(fmt.Sprintf(":%d", ServiceConstants.PortDefaultValude))
}
