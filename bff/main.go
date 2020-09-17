package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"light-up-backend/bff/impl"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"net/http"
)

func main() {
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	applicationConfig := common.GetApplicationConfig()
	config := applicationConfig.BffConfig()
	logger.WithField("service-name", config.ServiceName).Infoln("Starting service...")

	service := utils.CreateService(config.ServiceName)
	service.Init()

	r := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	n.Use(negronilogrus.NewCustomMiddleware(logrus.InfoLevel, &logrus.JSONFormatter{}, "negroni"))

	// Authentication.
	impl.RegisterLoginEndpoints(r, service.Client())
	// generic-api
	impl.RegisterGenericEndpoints(r, service.Client(), applicationConfig)

	go func() {
		port := config.Port
		logger.WithField("Port", port).Infoln("Starting Rest services now.")
		corsHandler := cors.AllowAll().Handler(r)
		handler := handlers.RecoveryHandler()(corsHandler)
		n.UseHandler(handler)
		err := http.ListenAndServe(config.Port, handler)
		if err != nil {
			logger.WithField("Error", err.Error()).Panicln("Could not start the server.")
		}
	}()
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
