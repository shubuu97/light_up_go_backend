package main

import (
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"light-up-backend/light-seeker-service/impl"
	"light-up-backend/light-seeker-service/proto"
)

func main() {
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	applicationConfigs := common.GetApplicationConfig()
	configs := applicationConfigs.LightSeekerServiceConfiguration()

	service := utils.CreateService(configs.ServiceName)
	repository := impl.NewLightSeekerRepository(ctx, configs)
	defer repository.Close()

	handler := impl.NewHandler(
		impl.NewLightSeekerService(
			repository,
			service.Client(),
		),
	)

	err := proto.RegisterLightSeekerServiceHandler(
		service.Server(),
		&handler,
	)

	if err != nil {
		logger.WithField("Error", err.Error()).Panicln("Could not register the handler.")
	}

	if err = service.Run(); err != nil {
		logger.Fatal(err)
	}
}
