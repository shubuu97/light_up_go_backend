package proto

import (
	"github.com/micro/go-micro/client"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

func CreateNewLightSeekerServiceClient(client client.Client) LightSeekerService {
	return NewLightSeekerService(common.LightSeekerServiceName, middleware.NewClientWrapper(client))
}
