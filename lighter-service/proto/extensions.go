package proto

import (
	"github.com/micro/go-micro/client"
	"light-up-backend/common"
	"light-up-backend/common/middleware"
)

func CreateNewLighterServiceClient(client client.Client) LighterService {
	return NewLighterService(common.LighterServiceName, middleware.NewClientWrapper(client))
}
