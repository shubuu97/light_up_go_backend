package main

import (
	"light-up-backend/authentication-service/impl"
	proto3 "light-up-backend/authentication-service/proto"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"light-up-backend/light-seeker-service/proto"
	proto2 "light-up-backend/lighter-service/proto"
)

func main() {
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	config := common.GetApplicationConfig().AuthenticationServiceConfig()
	service := utils.CreateService(config.ServiceName)

	service.Init()

	handler := impl.Handler{
		Service: &impl.Service{
			LightSeekerClient: proto.CreateNewLightSeekerServiceClient(service.Client()),
			LighterClient:     proto2.CreateNewLighterServiceClient(service.Client()),
		},
	}

	err := proto3.RegisterAuthenticationServiceHandler(
		service.Server(),
		&handler,
	)
	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not register the handler.")
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
