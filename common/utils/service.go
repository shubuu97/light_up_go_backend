package utils

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"light-up-backend/common/middleware"
	"time"
)

func init() {
	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	//customFormatter.FullTimestamp = true
}

func initializeLogging(serviceName string) {
	logLocation := fmt.Sprintf("/tmp/%s/%s-%s.log", serviceName, serviceName, time.Now().Format("2006-01-02 15:04:05"))
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logLocation,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     3,
		LocalTime:  true,
	}
	logrus.SetOutput(lumberjackLogger)
}

func CreateService(serviceName string) micro.Service {
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(middleware.Log),
		micro.WrapHandler(middleware.Instrument),
		//micro.WrapHandler(middleware.Validate),
		micro.RegisterInterval(time.Second*5),
		micro.RegisterTTL(time.Second*15),
	)
	service.Init()

	log.SetLogger(NewServiceLogger())

	initializeLogging(serviceName)
	return service
}
