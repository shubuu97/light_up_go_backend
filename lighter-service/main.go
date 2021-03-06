package main

import (
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
	"light-up-backend/lighter-service/impl"
	"light-up-backend/lighter-service/proto"
)

func main() {
	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	applicationConfigs := common.GetApplicationConfig()
	configs := applicationConfigs.LighterServiceConfiguration()
	service := utils.CreateService(configs.ServiceName)
	repository := impl.NewLighterRepository(ctx, configs)
	defer repository.Close()

	handler := impl.NewHandler(
		impl.NewLighterService(
			repository,
			service.Client(),
		),
	)

	err := proto.RegisterLighterServiceHandler(
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
