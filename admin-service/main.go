package main

import (
	"light-up-backend/admin-service/impl"
	"light-up-backend/admin-service/proto"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
	"light-up-backend/common/utils"
)

func main() {

	ctx := middleware.NewContext("main")
	logger := middleware.GetLogger(ctx)
	applicationConfigs := common.GetApplicationConfig()
	configs := applicationConfigs.AdminServiceConfig()
	service := utils.CreateService(configs.ServiceName)
	repository := impl.NewAdminRepository(ctx, configs)
	defer repository.Close()

	handler := impl.NewHandler(
		impl.NewAdminService(
			repository,
			service.Client(),
		),
	)

	err := proto.RegisterAdminServiceHandler(
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
